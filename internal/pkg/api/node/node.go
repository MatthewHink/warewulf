package node

import (
	"strconv"
	"strings"
	"github.com/hpcng/warewulf/internal/pkg/node"
	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
)

func NodeList(nodeNames []string) (nodeInfo []*wwapi.NodeInfo, err error) {
	// TODO: All logic from wwctl node list goes here.

	nodeDB, err := node.New()
	if err != nil {
		return
	}

	nodes, err := nodeDB.FindAllNodes()
	if err != nil {
		return
	}

	// Translate to the protobuf structure so wwapi can use this.
	// This is the same logic as for wwctl.
	for _, node := range node.FilterByName(nodes, nodeNames) {

		var ni wwapi.NodeInfo

		ni.NodeName = node.Id.Get()

		ni.Id.Source = node.Id.Source()
		ni.Id.Value = node.Id.Print()

		ni.Comment.Source = node.Comment.Source()
		ni.Comment.Value = node.Comment.Print()

		ni.Cluster.Source = node.ClusterName.Source()
		ni.Cluster.Value = node.ClusterName.Print()

		// source unused here
		ni.Profiles.Value = strings.Join(node.Profiles, ",")
		
		ni.Discoverable.Source = node.Discoverable.Source()
		ni.Discoverable.Value = strconv.FormatBool(node.Discoverable.PrintB())

		ni.Container.Source = node.ContainerName.Source()
		ni.Container.Value = node.ContainerName.Print()

		ni.Kernel.Source = node.KernelVersion.Source()
		ni.Kernel.Value = node.KernelVersion.Print()

		ni.KernelArgs.Source = node.KernelArgs.Source()
		ni.KernelArgs.Value = node.KernelArgs.Print()

		ni.SystemOverlay.Source = node.SystemOverlay.Source()
		ni.SystemOverlay.Value = node.SystemOverlay.Print()

		ni.RuntimeOverlay.Source = node.RuntimeOverlay.Source()
		ni.RuntimeOverlay.Value = node.RuntimeOverlay.Print()

		ni.Ipxe.Source = node.Ipxe.Source()
		ni.Ipxe.Value = node.Ipxe.Print()

		ni.Init.Source = node.Init.Source()
		ni.Init.Value = node.Init.Print()

		ni.Root.Source = node.Root.Source()
		ni.Root.Value = node.Root.Print()

		ni.AssetKey.Source = node.AssetKey.Source()
		ni.AssetKey.Value = node.AssetKey.Print()

		ni.IpmiIpaddr.Source = node.IpmiIpaddr.Source()
		ni.IpmiIpaddr.Value = node.IpmiIpaddr.Print()
		
		ni.IpmiNetmask.Source = node.IpmiNetmask.Source()
		ni.IpmiNetmask.Value = node.IpmiNetmask.Print()

		ni.IpmiPort.Source = node.IpmiPort.Source()
		ni.IpmiPort.Value = node.IpmiPort.Print()

		ni.IpmiGateway.Source = node.IpmiGateway.Source()
		ni.IpmiGateway.Value = node.IpmiGateway.Print()

		ni.IpmiUserName.Source = node.IpmiUserName.Source()
		ni.IpmiUserName.Value = node.IpmiUserName.Print()

		ni.IpmiInterface.Source = node.IpmiInterface.Source()
		ni.IpmiInterface.Value = node.IpmiInterface.Print()

		for keyname, keyvalue := range node.Tags{
			ni.Tags[keyname].Source = keyvalue.Source()
			ni.Tags[keyname].Value = keyvalue.Print()
		}

		for name, netdev := range node.NetDevs {
			ni.NetDevs[name].Device.Source = netdev.Device.Source()
			ni.NetDevs[name].Device.Value = netdev.Device.Print()

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
		}
		nodeInfo = append(nodeInfo, &ni)
	}
	return
} 

