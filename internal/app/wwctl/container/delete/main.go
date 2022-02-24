package delete

import (
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/container"
	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	cdp := &wwapiv1.ContainerDeleteParameter{
		ContainerNames: args,
	}
	return wwapi.ContainerDelete(cdp)
}
