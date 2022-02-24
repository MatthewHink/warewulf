package container

import (
	"github.com/hpcng/warewulf/internal/pkg/container"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
)

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
