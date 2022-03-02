package list

import (
	"fmt"

	"github.com/hpcng/warewulf/internal/pkg/api/container"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	containerInfo, err := container.ContainerList()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "%s\n", err)
		return
	}

	fmt.Printf("%-35s %-6s %-6s\n", "CONTAINER NAME", "BUILT", "NODES")
	for i := 0; i < len(containerInfo); i++ {
		fmt.Printf("%-35s %-6t %-6d\n",
			containerInfo[i].Name,
			containerInfo[i].Built,
			containerInfo[i].NodeCount)
	}
	return
}
