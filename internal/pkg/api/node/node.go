package node

import (
	"fmt"
	"os"
	"strconv"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/warewulfd"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/hpcng/warewulf/pkg/hostlist"
	"github.com/pkg/errors"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
)

// TODO: null checks on all pointer parameters

func NodeAdd(nap *wwapi.NodeAddParameter) (err error) {
	var count uint
	nodeDB, err := node.New()
	if err != nil {
		return errors.Wrap(err, "failed to open node database")
	}

	//node_args := hostlist.Expand(nap.NodeNames.NodeNames)
	node_args := hostlist.Expand(nap.NodeNames)

	for _, a := range node_args {
		n, err := nodeDB.AddNode(a)
		if err != nil {
			return errors.Wrap(err, "failed to add node")
		}
		wwlog.Printf(wwlog.INFO, "Added node: %s\n", a)

		if nap.Cluster != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting cluster name to: %s\n", n.Id.Get(), nap.Cluster)
			n.ClusterName.Set(nap.Cluster)
			err := nodeDB.NodeUpdate(n)
			if err != nil {
				return errors.Wrap(err, "failed to update node")
			}
		}

		if nap.Netdev != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				var netdev node.NetDevEntry
				n.NetDevs[nap.Netname] = &netdev
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting Device to: %s\n", n.Id.Get(), nap.Netname, nap.Netdev)

			n.NetDevs[nap.Netname].Device.Set(nap.Netdev)
			n.NetDevs[nap.Netname].OnBoot.SetB(true)
		}

		if nap.Ipaddr != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			NewIpaddr := util.IncrementIPv4(nap.Ipaddr, count)

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				var netdev node.NetDevEntry
				n.NetDevs[nap.Netname] = &netdev
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting Ipaddr to: %s\n", n.Id.Get(), nap.Netname, NewIpaddr)

			n.NetDevs[nap.Netname].Ipaddr.Set(NewIpaddr)
			n.NetDevs[nap.Netname].OnBoot.SetB(true)
		}

		if nap.Netmask != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				return errors.New("network device does not exist: " + nap.Netname)
			}
			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting netmask to: %s\n", n.Id.Get(), nap.Netname, nap.Netmask)

			n.NetDevs[nap.Netname].Netmask.Set(nap.Netmask)
		}

		if nap.Gateway != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				return errors.New("network device does not exist: " + nap.Netname)
			}
			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting gateway to: %s\n", n.Id.Get(), nap.Netname, nap.Gateway)

			n.NetDevs[nap.Netname].Gateway.Set(nap.Gateway)
		}

		if nap.Hwaddr != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				return errors.New("network device does not exist: " + nap.Netname)
			}
			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting HW address to: %s\n", n.Id.Get(), nap.Netname, nap.Hwaddr)

			n.NetDevs[nap.Netname].Hwaddr.Set(nap.Hwaddr)
			n.NetDevs[nap.Netname].OnBoot.SetB(true)
		}

		if nap.Type != "" {
			if nap.Netname == "" {
				return errors.New("you must include the '--netname' option")
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				return errors.New("network device does not exist: " + nap.Netname)
			}
			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting Type to: %s\n", n.Id.Get(), nap.Netname, nap.Type)

			n.NetDevs[nap.Netname].Type.Set(nap.Type)
		}

		if nap.Discoverable {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting node to discoverable\n", n.Id.Get())

			n.Discoverable.SetB(true)
		}

		err = nodeDB.NodeUpdate(n)
		if err != nil {
			return errors.Wrap(err, "failed to update nodedb")
		}

		count++
	} // end for
	
	err = nodeDB.Persist()
	if err != nil {
		return errors.Wrap(err, "failed to persist new node")
	}

	err = warewulfd.DaemonReload()
	if err != nil {
		return errors.Wrap(err, "failed to reload warewulf daemon")
	}
	return
}

func NodeDelete(ndp *wwapi.NodeDeleteParameter) (err error) {
	var count int
	var nodeList []node.NodeInfo

	nodeDB, err := node.New()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Failed to open node database: %s\n", err)
		return
	}

	nodes, err := nodeDB.FindAllNodes()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not get node list: %s\n", err)
		return
	}

	//node_args := hostlist.Expand(ndp.NodeNames.NodeNames)
	node_args := hostlist.Expand(ndp.NodeNames)

	for _, r := range node_args {
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
		return
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

	err = nodeDB.Persist()
	if err != nil {
		return errors.Wrap(err, "failed to persist nodedb")
	}

	err = warewulfd.DaemonReload()
	if err != nil {
		return errors.Wrap(err, "failed to reload warewulf daemon")
	}

	return
}


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