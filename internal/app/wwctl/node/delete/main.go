package delete

import (
	"fmt"
	"os"

	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/node" // TODO: Rename package/file to apinode? May be easier.
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/hpcng/warewulf/pkg/hostlist"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	// For a blocking prompt, expand the node list first.
	// We do this again at the API layer, but an API should not have a blocking prompt.
	// TODO: Break this up into:
	// - NodeDeleteParamCheck(consoleOutputOnOff)
	// - NodeDeletePrompt.
	if !SetYes {
		var count int
		var nodeList []node.NodeInfo

		nodeDB, err := node.New()
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "Failed to open node database: %s\n", err)
			return err
		}

		nodes, err := nodeDB.FindAllNodes()
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "Could not get node list: %s\n", err)
			return err
		}

		args = hostlist.Expand(args)

		for _, r := range args {
			var match bool
			for _, n := range nodes {
				if n.Id.Get() == r {
					nodeList = append(nodeList, n)
					match = true
				}
			}

			if !match {
				fmt.Fprintf(os.Stderr, "ERROR: No match for node: %s\n", r)
			}
		}

		if len(nodeList) == 0 {
			fmt.Printf("No nodes found\n")
			return nil
		}

		for _, n := range nodeList {
			err := nodeDB.DelNode(n.Id.Get())
			if err != nil {
				wwlog.Printf(wwlog.ERROR, "%s\n", err)
			} else {
				count++
				fmt.Printf("Deleting node: %s\n", n.Id.Print())
			}
		}
		q := fmt.Sprintf("Are you sure you want to delete %d nodes(s)", count)

		prompt := promptui.Prompt{
			Label:     q,
			IsConfirm: true,
		}

		result, _ := prompt.Run()

		if !(result == "y" || result == "yes") {
			return nil
		}
	}

	// Call the API to delete the nodes.
	ndp := wwapiv1.NodeDeleteParameter{
		Force: SetForce,
		NodeNames: args,
	}
	return wwapi.NodeDelete(&ndp)
}
