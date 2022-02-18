package list

import (
	"fmt"
	//"os"
	"sort"
	"strings"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/node"
	//"github.com/hpcng/warewulf/internal/pkg/node"
	//"github.com/hpcng/warewulf/internal/pkg/wwlog"
	//"github.com/hpcng/warewulf/pkg/hostlist"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	nodeInfo, err := wwapi.NodeList(args)
	if err != nil {
		return
	}
	//fmt.Printf("nodeInfo: %#v\n", nodeInfo)

	if ShowAll {
		for i := 0; i < len(nodeInfo); i++ {
			ni := nodeInfo[i]
			nodeName := ni.NodeName

			//fmt.Printf("ni: %#v\n", ni)

			fmt.Printf("################################################################################\n")
			fmt.Printf("%-20s %-18s %-12s %s\n", "NODE", "FIELD", "PROFILE", "VALUE")
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Id", ni.Id.Source, ni.Id.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Comment", ni.Comment.Source, ni.Comment.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Cluster", ni.Cluster.Source, ni.Cluster.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Profiles", "--", ni.Profiles.Value)

			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Discoverable", ni.Discoverable.Source, ni.Discoverable.Value)

			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Container", ni.Container.Source, ni.Container.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Kernel", ni.Kernel.Source, ni.Kernel.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "KernelArgs", ni.KernelArgs.Source, ni.KernelArgs.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "SystemOverlay", ni.SystemOverlay.Source, ni.SystemOverlay.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "RuntimeOverlay", ni.RuntimeOverlay.Source, ni.RuntimeOverlay.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Ipxe", ni.Ipxe.Source, ni.Ipxe.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Init", ni.Init.Source, ni.Init.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Root", ni.Root.Source, ni.Root.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "AssetKey", ni.AssetKey.Source, ni.AssetKey.Value)

			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiIpaddr", ni.IpmiIpaddr.Source, ni.IpmiIpaddr.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiNetmask", ni.IpmiNetmask.Source, ni.IpmiNetmask.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiPort", ni.IpmiPort.Source, ni.IpmiPort.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiGateway", ni.IpmiGateway.Source, ni.IpmiGateway.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiUserName", ni.IpmiUserName.Source, ni.IpmiUserName.Value)
			fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "IpmiInterface", ni.IpmiInterface.Source, ni.IpmiInterface.Value)

			for keyname, key := range ni.Tags {
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, "Tag["+keyname+"]", key.Source, key.Value)
			}

			for name, netdev := range ni.NetDevs {
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":DEVICE", netdev.Device.Source, netdev.Device.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":HWADDR", netdev.Hwaddr.Source, netdev.Hwaddr.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":IPADDR", netdev.Ipaddr.Source, netdev.Ipaddr.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":NETMASK", netdev.Netmask.Source, netdev.Netmask.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":GATEWAY", netdev.Gateway.Source, netdev.Gateway.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":TYPE", netdev.Type.Source, netdev.Type.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":ONBOOT", netdev.Onboot.Source, netdev.Onboot.Value)
				fmt.Printf("%-20s %-18s %-12s %s\n", nodeName, name+":DEFAULT", netdev.Default.Source, netdev.Default.Value)
			}
		}
	} else if ShowNet {
		fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", "NODE NAME", "DEVICE", "HWADDR", "IPADDR", "GATEWAY")
		fmt.Println(strings.Repeat("=", 80))

		for i := 0; i < len(nodeInfo); i++ {
			ni := nodeInfo[i]
			nodeName := ni.NodeName

			if len(ni.NetDevs) > 0 {
				for name, dev := range ni.NetDevs {
					fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", nodeName, name, dev.Hwaddr.Value, dev.Ipaddr.Value, dev.Gateway.Value)
				}
			} else {
				fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", nodeName, "--", "--", "--", "--")
			}
		}
	} else if ShowIpmi {
		fmt.Printf("%-22s %-16s %-10s %-20s %-20s %-14s\n", "NODE NAME", "IPMI IPADDR", "IPMI PORT", "IPMI USERNAME", "IPMI PASSWORD", "IPMI INTERFACE")
		fmt.Println(strings.Repeat("=", 108))

		for i := 0; i < len(nodeInfo); i++ {
			ni := nodeInfo[i]
			nodeName := ni.NodeName
			fmt.Printf("%-22s %-16s %-10s %-20s %-20s %-14s\n", nodeName, ni.IpmiIpaddr.Value, ni.IpmiPort.Value, ni.IpmiUserName.Value, ni.IpmiPassword.Value, ni.IpmiInterface.Value)
		}

	} else if ShowLong {
		fmt.Printf("%-22s %-26s %-35s %s\n", "NODE NAME", "KERNEL", "CONTAINER", "OVERLAYS (S/R)")
		fmt.Println(strings.Repeat("=", 120))

		for i := 0; i < len(nodeInfo); i++ {
			ni := nodeInfo[i]
			nodeName := ni.NodeName
			fmt.Printf("%-22s %-26s %-35s %s\n", nodeName, ni.Kernel.Value, ni.Container.Value, ni.SystemOverlay.Value+"/"+ni.RuntimeOverlay.Value)
		}

	} else {
		fmt.Printf("%-22s %-26s %s\n", "NODE NAME", "PROFILES", "NETWORK")
		fmt.Println(strings.Repeat("=", 80))

		for i := 0; i < len(nodeInfo); i++ {
			ni := nodeInfo[i]
			//nodeName := ni.NodeName
			var netdevs []string
			if len(ni.NetDevs) > 0 {
				for name, dev := range ni.NetDevs {
					netdevs = append(netdevs, fmt.Sprintf("%s:%s", name, dev.Ipaddr.Value))
				}
			}
			sort.Strings(netdevs)

			// TODO: Fix this.
			//fmt.Printf("%-22s %-26s %s\n", nodeName, strings.Join(ni.Profiles.Value, ","), strings.Join(netdevs, ", "))
		}

	}

	/*
	if ShowAll {
		for i :=  0; i < nodeInfo.Len(); i++ {

			fmt.Printf("################################################################################\n")
			fmt.Printf("%-20s %-18s %-12s %s\n", "NODE", "FIELD", "PROFILE", "VALUE")
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Id", node.Id.Source(), node.Id.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Comment", node.Comment.Source(), node.Comment.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Cluster", node.ClusterName.Source(), node.ClusterName.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Profiles", "--", strings.Join(node.Profiles, ","))

			fmt.Printf("%-20s %-18s %-12s %t\n", node.Id.Get(), "Discoverable", node.Discoverable.Source(), node.Discoverable.PrintB())

			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Container", node.ContainerName.Source(), node.ContainerName.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Kernel", node.KernelVersion.Source(), node.KernelVersion.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "KernelArgs", node.KernelArgs.Source(), node.KernelArgs.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "SystemOverlay", node.SystemOverlay.Source(), node.SystemOverlay.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "RuntimeOverlay", node.RuntimeOverlay.Source(), node.RuntimeOverlay.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Ipxe", node.Ipxe.Source(), node.Ipxe.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Init", node.Init.Source(), node.Init.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Root", node.Root.Source(), node.Root.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "AssetKey", node.AssetKey.Source(), node.AssetKey.Print())

			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiIpaddr", node.IpmiIpaddr.Source(), node.IpmiIpaddr.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiNetmask", node.IpmiNetmask.Source(), node.IpmiNetmask.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiPort", node.IpmiPort.Source(), node.IpmiPort.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiGateway", node.IpmiGateway.Source(), node.IpmiGateway.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiUserName", node.IpmiUserName.Source(), node.IpmiUserName.Print())
			fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "IpmiInterface", node.IpmiInterface.Source(), node.IpmiInterface.Print())

			for keyname, key := range node.Tags {
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), "Tag["+keyname+"]", key.Source(), key.Print())
			}

			for name, netdev := range node.NetDevs {
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":DEVICE", netdev.Device.Source(), netdev.Device.Print())
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":HWADDR", netdev.Hwaddr.Source(), netdev.Hwaddr.Print())
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":IPADDR", netdev.Ipaddr.Source(), netdev.Ipaddr.Print())
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":NETMASK", netdev.Netmask.Source(), netdev.Netmask.Print())
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":GATEWAY", netdev.Gateway.Source(), netdev.Gateway.Print())
				fmt.Printf("%-20s %-18s %-12s %s\n", node.Id.Get(), name+":TYPE", netdev.Type.Source(), netdev.Type.Print())
				fmt.Printf("%-20s %-18s %-12s %t\n", node.Id.Get(), name+":ONBOOT", netdev.OnBoot.Source(), netdev.OnBoot.PrintB())
				fmt.Printf("%-20s %-18s %-12s %t\n", node.Id.Get(), name+":DEFAULT", netdev.Default.Source(), netdev.Default.PrintB())
			}

		}

	} else if ShowNet {
		fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", "NODE NAME", "DEVICE", "HWADDR", "IPADDR", "GATEWAY")
		fmt.Println(strings.Repeat("=", 80))

		for _, node := range node.FilterByName(nodes, args) {
			if len(node.NetDevs) > 0 {
				for name, dev := range node.NetDevs {
					fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", node.Id.Get(), name, dev.Hwaddr.Print(), dev.Ipaddr.Print(), dev.Gateway.Print())
				}
			} else {
				fmt.Printf("%-22s %-6s %-18s %-15s %-15s\n", node.Id.Get(), "--", "--", "--", "--")
			}
		}

	} else if ShowIpmi {
		fmt.Printf("%-22s %-16s %-10s %-20s %-20s %-14s\n", "NODE NAME", "IPMI IPADDR", "IPMI PORT", "IPMI USERNAME", "IPMI PASSWORD", "IPMI INTERFACE")
		fmt.Println(strings.Repeat("=", 108))

		for _, node := range node.FilterByName(nodes, args) {
			fmt.Printf("%-22s %-16s %-10s %-20s %-20s %-14s\n", node.Id.Get(), node.IpmiIpaddr.Print(), node.IpmiPort.Print(), node.IpmiUserName.Print(), node.IpmiPassword.Print(), node.IpmiInterface.Print())
		}

	} else if ShowLong {
		fmt.Printf("%-22s %-26s %-35s %s\n", "NODE NAME", "KERNEL", "CONTAINER", "OVERLAYS (S/R)")
		fmt.Println(strings.Repeat("=", 120))

		for _, node := range node.FilterByName(nodes, args) {
			fmt.Printf("%-22s %-26s %-35s %s\n", node.Id.Get(), node.KernelVersion.Print(), node.ContainerName.Print(), node.SystemOverlay.Print()+"/"+node.RuntimeOverlay.Print())
		}

	} else {
		fmt.Printf("%-22s %-26s %s\n", "NODE NAME", "PROFILES", "NETWORK")
		fmt.Println(strings.Repeat("=", 80))

		for _, node := range node.FilterByName(nodes, args) {
			var netdevs []string
			if len(node.NetDevs) > 0 {
				for name, dev := range node.NetDevs {
					netdevs = append(netdevs, fmt.Sprintf("%s:%s", name, dev.Ipaddr.Print()))
				}
			}
			sort.Strings(netdevs)

			fmt.Printf("%-22s %-26s %s\n", node.Id.Get(), strings.Join(node.Profiles, ","), strings.Join(netdevs, ", "))
		}

	}
	return nil
	*/
	return
}
