package add

import (
	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/node"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {

	nap := wwapiv1.NodeAddParameter{
		Cluster: SetClusterName,
		Discoverable: SetDiscoverable,
		Gateway: SetGateway,
		Hwaddr: SetHwaddr,
		Ipaddr: SetIpaddr,
		Netdev: SetNetDev,
		Netmask: SetNetmask,
		Netname: SetNetName,
		Type: SetType,
		NodeNames: args,
	}
	return wwapi.NodeAdd(&nap)
}