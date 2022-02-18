package node

import (
	//"fmt"
	"strconv"
	"strings"
	"github.com/hpcng/warewulf/internal/pkg/node"
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

	// Translate to the protobuf structure so wwapi can use this across the wire.
	// This is the same logic as was in wwctl.
	for _, node := range node.FilterByName(nodes, nodeNames) {
		//fmt.Printf("node: %#v\n", node)
		//fmt.Printf("node.Id: %#v\n", node.Id)


		var ni wwapi.NodeInfo // TODO: This is uninitialized. It has no Id.

		ni.NodeName = node.Id.Get()

		//fmt.Printf("node Id %v\n", node.Id)
		//fmt.Printf("node Id.Source %v\n", node.Id.Source())

		//fmt.Printf("ni.Id: %#v\n", ni.Id) // TODO: This is your problem.

		//ni.Id = node.Id
		//ni.Id = &wwapi.NodeInfo.Id{}
		ni.Id = &wwapi.NodeField{
			Source: node.Id.Source(),
			Value: node.Id.Print(),
		}

		//ni.Id.Source = node.Id.Source()
		//ni.Id.Value = node.Id.Print()

		ni.Comment = &wwapi.NodeField{
			Source: node.Comment.Source(),
			Value: node.Comment.Print(),
		}

		//ni.Comment.Source = node.Comment.Source()
		//ni.Comment.Value = node.Comment.Print()

		ni.Cluster = &wwapi.NodeField{
			Source: node.ClusterName.Source(),
			Value: node.ClusterName.Print(),
		}
		//ni.Cluster.Source = node.ClusterName.Source()
		//ni.Cluster.Value = node.ClusterName.Print()

		// source unused here
		ni.Profiles = &wwapi.NodeField{
			Value: strings.Join(node.Profiles, ","),
		}
		//ni.Profiles.Value = strings.Join(node.Profiles, ",")
		
		ni.Discoverable = &wwapi.NodeField {
			Source: node.Discoverable.Source(),
			Value: strconv.FormatBool(node.Discoverable.PrintB()),
		}
		//ni.Discoverable.Source = node.Discoverable.Source()
		//ni.Discoverable.Value = strconv.FormatBool(node.Discoverable.PrintB())

		ni.Container = &wwapi.NodeField {
			Source: node.ContainerName.Source(),
			Value: node.ContainerName.Print(),
		}
		//ni.Container.Source = node.ContainerName.Source()
		//ni.Container.Value = node.ContainerName.Print()

		ni.Kernel = &wwapi.NodeField{
			Source: node.KernelVersion.Source(),
			Value: node.KernelVersion.Print(),
		}
		//ni.Kernel.Source = node.KernelVersion.Source()
		//ni.Kernel.Value = node.KernelVersion.Print()

		ni.KernelArgs = &wwapi.NodeField{
			Source: node.KernelArgs.Source(),
			Value: node.KernelArgs.Print(),
		}
		//ni.KernelArgs.Source = node.KernelArgs.Source()
		//ni.KernelArgs.Value = node.KernelArgs.Print()


		ni.SystemOverlay = &wwapi.NodeField{
			Source: node.SystemOverlay.Source(),
			Value: node.SystemOverlay.Print(),
		}

		//ni.SystemOverlay.Source = node.SystemOverlay.Source()
		//ni.SystemOverlay.Value = node.SystemOverlay.Print()

		ni.RuntimeOverlay = &wwapi.NodeField{
			Source: node.RuntimeOverlay.Source(),
			Value: node.RuntimeOverlay.Print(),
		}
		//ni.RuntimeOverlay.Source = node.RuntimeOverlay.Source()
		//ni.RuntimeOverlay.Value = node.RuntimeOverlay.Print()

		ni.Ipxe = &wwapi.NodeField{
			Source: node.Ipxe.Source(),
			Value: node.Ipxe.Print(),
		}
		//ni.Ipxe.Source = node.Ipxe.Source()
		//ni.Ipxe.Value = node.Ipxe.Print()

		ni.Init = &wwapi.NodeField{
			Source: node.Init.Source(),
			Value: node.Init.Print(),
		}
		//ni.Init.Source = node.Init.Source()
		//ni.Init.Value = node.Init.Print()

		ni.Root = &wwapi.NodeField{
			Source: node.Root.Source(),
			Value: node.Root.Print(),
		}
		//ni.Root.Source = node.Root.Source()
		//ni.Root.Value = node.Root.Print()

		ni.AssetKey = &wwapi.NodeField{
			Source: node.AssetKey.Source(),
			Value: node.AssetKey.Print(),
		}
		//ni.AssetKey.Source = node.AssetKey.Source()
		//ni.AssetKey.Value = node.AssetKey.Print()

		ni.IpmiIpaddr = &wwapi.NodeField{
			Source: node.IpmiIpaddr.Source(),
			Value: node.IpmiIpaddr.Print(),
		}

		//ni.IpmiIpaddr.Source = node.IpmiIpaddr.Source()
		//ni.IpmiIpaddr.Value = node.IpmiIpaddr.Print()
		
		ni.IpmiNetmask = &wwapi.NodeField{
			Source: node.IpmiNetmask.Source(),
			Value: node.IpmiNetmask.Print(),
		}
		//ni.IpmiNetmask.Source = node.IpmiNetmask.Source()
		//ni.IpmiNetmask.Value = node.IpmiNetmask.Print()

		ni.IpmiPort = &wwapi.NodeField{
			Source: node.IpmiPort.Source(),
			Value: node.IpmiPort.Print(),
		}
		//ni.IpmiPort.Source = node.IpmiPort.Source()
		//ni.IpmiPort.Value = node.IpmiPort.Print()

		ni.IpmiGateway = &wwapi.NodeField{
			Source: node.IpmiGateway.Source(),
			Value: node.IpmiGateway.Print(),
		}
		//ni.IpmiGateway.Source = node.IpmiGateway.Source()
		//ni.IpmiGateway.Value = node.IpmiGateway.Print()

		ni.IpmiUserName = &wwapi.NodeField{
			Source: node.IpmiUserName.Source(),
			Value: node.IpmiUserName.Print(),
		}
		//ni.IpmiUserName.Source = node.IpmiUserName.Source()
		//ni.IpmiUserName.Value = node.IpmiUserName.Print()

		ni.IpmiInterface = &wwapi.NodeField{
			Source: node.IpmiInterface.Source(),
			Value: node.IpmiInterface.Print(),
		}
		//ni.IpmiInterface.Source = node.IpmiInterface.Source()
		//ni.IpmiInterface.Value = node.IpmiInterface.Print()

		for keyname, keyvalue := range node.Tags{
			ni.Tags[keyname].Source = keyvalue.Source()
			ni.Tags[keyname].Value = keyvalue.Print()
		}

		ni.NetDevs = map[string]*wwapi.NetDev{}
		for name, netdev := range node.NetDevs {
			//ni.NetDevs[name].Device = &wwapi.NetDev{
			//	Device: &wwapi.NodeField{
			//		Source: netdev.Device.Source(),
			//		Value: netdev.Device.Print(),
			//	},
			//}
			//ni.NetDevs = map[string]*wwapi.NetDev{}


			ni.NetDevs[name] = &wwapi.NetDev{
				Device: &wwapi.NodeField{
					Source: netdev.Device.Source(),
					Value: netdev.Device.Print(),
				},
				Hwaddr: &wwapi.NodeField{
					Source: netdev.Hwaddr.Source(),
					Value: netdev.Hwaddr.Print(),
				},
				Ipaddr: &wwapi.NodeField{
					Source: netdev.Ipaddr.Source(),
					Value: netdev.Ipaddr.Print(),
				},
				Netmask: &wwapi.NodeField{
					Source: netdev.Netmask.Source(),
					Value: netdev.Netmask.Print(),
				},
				Gateway: &wwapi.NodeField{
					Source: netdev.Gateway.Source(),
					Value: netdev.Gateway.Print(),
				},
				Type: &wwapi.NodeField{
					Source: netdev.Type.Source(),
					Value: netdev.Type.Print(),
				},
				Onboot: &wwapi.NodeField{
					Source: netdev.OnBoot.Source(),
					Value: strconv.FormatBool(netdev.OnBoot.PrintB()),
				},				
				Default: &wwapi.NodeField{
					Source: netdev.Default.Source(),
					Value: strconv.FormatBool(netdev.Default.PrintB()),
				},

			}

			//ni.NetDevs[name].Device = &wwapi.NodeField{
			//	Source: netdev.Device.Source(),
			//	Value: netdev.Device.Print(),
			//}

			//ni.NetDevs[name].Device.Source = netdev.Device.Source()
			//ni.NetDevs[name].Device.Value = netdev.Device.Print()
/*
			ni.NetDevs[name].Hwaddr.Source = netdev.Hwaddr.Source()
			ni.NetDevs[name].Hwaddr.Value = netdev.Hwaddr.Print()

			ni.NetDevs[name].Ipaddr.Source = netdev.Ipaddr.Source()
			ni.NetDevs[name].Ipaddr.Value = netdev.Ipaddr.Print()

			ni.NetDevs[name].Netmask.Source = netdev.Netmask.Source()
			ni.NetDevs[name].Netmask.Value = netdev.Netmask.Print()

			ni.NetDevs[name].Gateway.Source = netdev.Gateway.Source()
			ni.NetDevs[name].Gateway.Value = netdev.Gateway.Print()

			ni.NetDevs[name].Type.Source = netdev.Type.Source()
			ni.NetDevs[name].Type.Value = netdev.Type.Print()

			ni.NetDevs[name].Onboot.Source = netdev.OnBoot.Source()
			ni.NetDevs[name].Onboot.Value = strconv.FormatBool(netdev.OnBoot.PrintB())

			ni.NetDevs[name].Default.Source = netdev.Default.Source()
			ni.NetDevs[name].Default.Value = strconv.FormatBool(netdev.Default.PrintB())
*/
		}
		
		nodeInfo = append(nodeInfo, &ni)
	}
	return
} 

