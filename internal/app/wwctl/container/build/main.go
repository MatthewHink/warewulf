package build

import (
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/container"
	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {

	cbp := &wwapiv1.ContainerBuildParameter{
		ContainerNames: args,
		Force: BuildForce,
		All: BuildAll,
		Default: SetDefault,
	}
	return wwapi.ContainerBuild(cbp)
}
