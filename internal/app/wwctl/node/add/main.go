package add

import (
	//"github.com/hpcng/warewulf/internal/pkg/node"
	//"github.com/hpcng/warewulf/internal/pkg/util"
	//"github.com/hpcng/warewulf/internal/pkg/warewulfd"
	//"github.com/hpcng/warewulf/internal/pkg/wwlog"
	//"github.com/hpcng/warewulf/pkg/hostlist"
	//"github.com/pkg/errors"
	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/spf13/cobra"	
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/node"
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
		NodeNames: &wwapiv1.NodeNames{
			NodeNames: args,
		},
	}
	return wwapi.NodeAdd(&nap)
}
