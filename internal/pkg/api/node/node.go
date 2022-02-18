package node

import (
	"strconv"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/pkg/hostlist"
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
)

func NodeList(nodeNames []string) (nodeInfo []*wwapi.NodeInfo, err error) {

	nodeDB, err := node.New()
	if err != nil {
		return
	}

	nodes, err := nodeDB.FindAllNodes()
	if err != nil {
		return
	}

	nodeNames = hostlist.Expand(nodeNames)

	// Translate to the protobuf structure so wwapi can use this across the wire.
	// This is the same logic as was in wwctl.
	for _, node := range node.FilterByName(nodes, nodeNames) {

		var ni wwapi.NodeInfo
		
		ni.NodeName = node.Id.Get()

		ni.Id = &wwapi.NodeField{
			Source: node.Id.Source(),
			Value: node.Id.Get(),
			Print: node.Id.Print(),
		}

		ni.Comment = &wwapi.NodeField{
			Source: node.Comment.Source(),
			Value: node.Comment.Get(),
			Print: node.Comment.Print(),
		}

		ni.Cluster = &wwapi.NodeField{
			Source: node.ClusterName.Source(),
			Value: node.ClusterName.Get(),
			Print: node.ClusterName.Print(),
		}

		// source unused here
		ni.Profiles = node.Profiles
		
		ni.Discoverable = &wwapi.NodeField {
			Source: node.Discoverable.Source(),
			Value: strconv.FormatBool(node.Discoverable.GetB()),
			Print: strconv.FormatBool(node.Discoverable.PrintB()),
		}

		ni.Container = &wwapi.NodeField {
			Source: node.ContainerName.Source(),
			Value: node.ContainerName.Get(),
			Print: node.ContainerName.Print(),
		}

		ni.Kernel = &wwapi.NodeField{
			Source: node.KernelVersion.Source(),
			Value: node.KernelVersion.Get(),
			Print: node.KernelVersion.Print(),
		}

		ni.KernelArgs = &wwapi.NodeField{
			Source: node.KernelArgs.Source(),
			Value: node.KernelArgs.Get(),
			Print: node.KernelArgs.Print(),
		}

		ni.SystemOverlay = &wwapi.NodeField{
			Source: node.SystemOverlay.Source(),
			Value: node.SystemOverlay.Get(),
			Print: node.SystemOverlay.Print(),
		}

		ni.RuntimeOverlay = &wwapi.NodeField{
			Source: node.RuntimeOverlay.Source(),
			Value: node.RuntimeOverlay.Get(),
			Print: node.RuntimeOverlay.Print(),
		}

		ni.Ipxe = &wwapi.NodeField{
			Source: node.Ipxe.Source(),
			Value: node.Ipxe.Get(),
			Print: node.Ipxe.Print(),
		}

		ni.Init = &wwapi.NodeField{
			Source: node.Init.Source(),
			Value: node.Init.Get(),
			Print: node.Init.Print(),
		}

		ni.Root = &wwapi.NodeField{
			Source: node.Root.Source(),
			Value: node.Root.Get(),
			Print: node.Root.Print(),
		}

		ni.AssetKey = &wwapi.NodeField{
			Source: node.AssetKey.Source(),
			Value: node.AssetKey.Get(),
			Print: node.AssetKey.Print(),
		}

		ni.IpmiIpaddr = &wwapi.NodeField{
			Source: node.IpmiIpaddr.Source(),
			Value: node.IpmiIpaddr.Get(),
			Print: node.IpmiIpaddr.Print(),
		}

		ni.IpmiNetmask = &wwapi.NodeField{
			Source: node.IpmiNetmask.Source(),
			Value: node.IpmiNetmask.Get(),
			Print: node.IpmiNetmask.Print(),
		}

		ni.IpmiPort = &wwapi.NodeField{
			Source: node.IpmiPort.Source(),
			Value: node.IpmiPort.Get(),
			Print: node.IpmiPort.Print(),
		}

		ni.IpmiGateway = &wwapi.NodeField{
			Source: node.IpmiGateway.Source(),
			Value: node.IpmiGateway.Get(),
			Print: node.IpmiGateway.Print(),
		}

		ni.IpmiUserName = &wwapi.NodeField{
			Source: node.IpmiUserName.Source(),
			Value: node.IpmiUserName.Get(),
			Print: node.IpmiUserName.Print(),
		}

		ni.IpmiPassword = &wwapi.NodeField{
			Source: node.IpmiPassword.Source(),
			Value: node.IpmiPassword.Get(),
			Print: node.IpmiPassword.Print(),
		}

		ni.IpmiInterface = &wwapi.NodeField{
			Source: node.IpmiInterface.Source(),
			Value: node.IpmiInterface.Get(),
			Print: node.IpmiInterface.Print(),
		}

		for keyname, keyvalue := range node.Tags{
			ni.Tags[keyname].Source = keyvalue.Source()
			ni.Tags[keyname].Value = keyvalue.Get()
			ni.Tags[keyname].Print = keyvalue.Print()
		}

		ni.NetDevs = map[string]*wwapi.NetDev{}
		for name, netdev := range node.NetDevs {

			ni.NetDevs[name] = &wwapi.NetDev{
				Device: &wwapi.NodeField{
					Source: netdev.Device.Source(),
					Value: netdev.Device.Get(),
					Print: netdev.Device.Print(),
				},
				Hwaddr: &wwapi.NodeField{
					Source: netdev.Hwaddr.Source(),
					Value: netdev.Hwaddr.Get(),
					Print: netdev.Hwaddr.Print(),
				},
				Ipaddr: &wwapi.NodeField{
					Source: netdev.Ipaddr.Source(),
					Value: netdev.Ipaddr.Get(),
					Print: netdev.Ipaddr.Print(),
				},
				Netmask: &wwapi.NodeField{
					Source: netdev.Netmask.Source(),
					Value: netdev.Netmask.Get(),
					Print: netdev.Netmask.Print(),
				},
				Gateway: &wwapi.NodeField{
					Source: netdev.Gateway.Source(),
					Value: netdev.Gateway.Get(),
					Print: netdev.Gateway.Print(),
				},
				Type: &wwapi.NodeField{
					Source: netdev.Type.Source(),
					Value: netdev.Type.Get(),
					Print: netdev.Type.Print(),
				},
				Onboot: &wwapi.NodeField{
					Source: netdev.OnBoot.Source(),
					Value: strconv.FormatBool(netdev.OnBoot.GetB()),
					Print: strconv.FormatBool(netdev.OnBoot.PrintB()),
				},				
				Default: &wwapi.NodeField{
					Source: netdev.Default.Source(),
					Value: strconv.FormatBool(netdev.Default.GetB()),
					Print: strconv.FormatBool(netdev.Default.PrintB()),
				},

			}
		}
		nodeInfo = append(nodeInfo, &ni)
	}
	return
}