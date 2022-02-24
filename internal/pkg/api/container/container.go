package container

import (


	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/containers/image/v5/types"

	"github.com/hpcng/warewulf/internal/pkg/container"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/warewulfd"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/pkg/errors"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
)

func ContainerImport(cip *wwapi.ContainerImportParameter) (containerName string, err error) {

	if cip == nil {
		err = fmt.Errorf("NodeAddParameter is nil")
		return
	}

	if cip.Name == "" {
		name := path.Base(cip.Source)
		fmt.Printf("Setting VNFS name: %s\n", name)
		cip.Name = name
	}
	if !container.ValidName(cip.Name) {
		err = fmt.Errorf("VNFS name contains illegal characters: %s", cip.Name)
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}

	containerName = cip.Name
	fullPath := container.SourceDir(cip.Name)

	if util.IsDir(fullPath) {
		if cip.Force {
			fmt.Printf("Overwriting existing VNFS\n")
			err = os.RemoveAll(fullPath)
			if err != nil {
				wwlog.Printf(wwlog.ERROR, "%s\n", err)
				return
			}
		} else if cip.Update {
			fmt.Printf("Updating existing VNFS\n")
		} else {
			err = fmt.Errorf("VNFS Name exists, specify --force, --update, or choose a different name: %s", cip.Name)
			wwlog.Printf(wwlog.ERROR, "%s\n", err)
			return
		}
	} else if strings.HasPrefix(cip.Source, "docker://") || strings.HasPrefix(cip.Source, "docker-daemon://") {
		var sCtx *types.SystemContext
		sCtx, err = getSystemContext()
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "%s\n", err)
			// return was missing here. Was that deliberate?
		}

		err = container.ImportDocker(cip.Source, cip.Name, sCtx)
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "Could not import image: %s\n", err)
			_ = container.DeleteSource(cip.Name)
			return
		}
	} else if util.IsDir(cip.Source) {
		err = container.ImportDirectory(cip.Source, cip.Name)
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "Could not import image: %s\n", err)
			_ = container.DeleteSource(cip.Name)
			return
		}
	} else {
		err = fmt.Errorf("Invalid dir or uri: %s", cip.Source)
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}

	fmt.Printf("Updating the container's /etc/resolv.conf\n")
	err = util.CopyFile("/etc/resolv.conf", path.Join(container.RootFsDir(cip.Name), "/etc/resolv.conf"))
	if err != nil {
		wwlog.Printf(wwlog.WARN, "Could not copy /etc/resolv.conf into container: %s\n", err)
	}


	fmt.Printf("Building container: %s\n", cip.Name)
	err = container.Build(cip.Name, true)
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not build container %s: %s\n", cip.Name, err)
		return
	}

	if cip.Default {
		var nodeDB node.NodeYaml
		nodeDB, err = node.New()
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "Could not open node configuration: %s\n", err)
			return
		}

		//TODO: Don't loop through profiles, instead have a nodeDB function that goes directly to the map
		profiles, _ := nodeDB.FindAllProfiles()
		for _, profile := range profiles {
			wwlog.Printf(wwlog.DEBUG, "Looking for profile default: %s\n", profile.Id.Get())
			if profile.Id.Get() == "default" {
				wwlog.Printf(wwlog.DEBUG, "Found profile default, setting container name to: %s\n", cip.Name)
				profile.ContainerName.Set(cip.Name)
				err = nodeDB.ProfileUpdate(profile)
				if err != nil {
					err = errors.Wrap(err, "failed to update profile")
					return
				}
			}
		}
		// TODO: We need this in a function with a flock around it.
		// Also need to understand if the daemon restart is only to
		// reload the config or if there is something more.
		err = nodeDB.Persist()
		if err != nil {
			err = errors.Wrap(err, "failed to persist nodedb")
			return
		}

		fmt.Printf("Set default profile to container: %s\n", cip.Name)
		err = warewulfd.DaemonReload()
		if err != nil {
			err =  errors.Wrap(err, "failed to reload warewulf daemon")
			return
		}
	}
	return
}

func ContainerList() (containerInfo []*wwapi.ContainerInfo, err error) {
	var sources []string

	sources, err = container.ListSources()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}

	nodeDB, err := node.New()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}
	
	nodes, err := nodeDB.FindAllNodes()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}

	nodemap := make(map[string]int)
	for _, n := range nodes {
		nodemap[n.ContainerName.Get()]++
	}

	for _, source := range sources {
		image := container.ImageFile(source)

		if nodemap[source] == 0 {
			nodemap[source] = 0
		}
		containerInfo = append(containerInfo, &wwapi.ContainerInfo{
			Name: source,
			Built: util.IsFile(image),
			NodeCount: uint32(nodemap[source]),
		})
	}
	return
}

// Private helpers

func setOCICredentials(sCtx *types.SystemContext) error {
	username, userSet := os.LookupEnv("WAREWULF_OCI_USERNAME")
	password, passSet := os.LookupEnv("WAREWULF_OCI_PASSWORD")
	if userSet || passSet {
		if userSet && passSet {
			sCtx.DockerAuthConfig = &types.DockerAuthConfig{
				Username: username,
				Password: password,
			}
		} else {
			return fmt.Errorf("oci username and password env vars must be specified together")
		}
	}
	return nil
}

func setNoHTTPSOpts(sCtx *types.SystemContext) error {
	val, ok := os.LookupEnv("WAREWULF_OCI_NOHTTPS")
	if !ok {
		return nil
	}

	noHTTPS, err := strconv.ParseBool(val)
	if err != nil {
		return fmt.Errorf("while parsing insecure http option: %v", err)
	}

	// only set this if we want to disable, otherwise leave as undefined
	if noHTTPS {
		sCtx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(true)
	}
	sCtx.OCIInsecureSkipTLSVerify = noHTTPS

	return nil
}

func getSystemContext() (sCtx *types.SystemContext, err error) {
	sCtx = &types.SystemContext{}

	if err := setOCICredentials(sCtx); err != nil {
		return nil, err
	}

	if err := setNoHTTPSOpts(sCtx); err != nil {
		return nil, err
	}

	return sCtx, nil
}
