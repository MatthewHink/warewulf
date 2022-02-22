package node

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/hpcng/warewulf/internal/pkg/node"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/warewulfd"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/hpcng/warewulf/pkg/hostlist"
	"github.com/manifoldco/promptui"
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

	node_args := hostlist.Expand(nap.NodeNames)

	for _, a := range node_args {
		var n node.NodeInfo
		n, err = nodeDB.AddNode(a)
		if err != nil {
			return errors.Wrap(err, "failed to add node")
		}
		wwlog.Printf(wwlog.INFO, "Added node: %s\n", a)

		if nap.Cluster != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting cluster name to: %s\n", n.Id.Get(), nap.Cluster)
			n.ClusterName.Set(nap.Cluster)
			err = nodeDB.NodeUpdate(n)
			if err != nil {
				return errors.Wrap(err, "failed to update node")
			}
		}

		if nap.Netdev != "" {
			err = checkNetNameRequired(nap.Netname)
			if err != nil {
				return
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
			err = checkNetNameRequired(nap.Netname)
			if err != nil {
				return
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
			err = checkNetNameRequired(nap.Netname)
			if err != nil {
				return
			}

			if _, ok := n.NetDevs[nap.Netname]; !ok {
				return errors.New("network device does not exist: " + nap.Netname)
			}
			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting netmask to: %s\n", n.Id.Get(), nap.Netname, nap.Netmask)

			n.NetDevs[nap.Netname].Netmask.Set(nap.Netmask)
		}

		if nap.Gateway != "" {
			err = checkNetNameRequired(nap.Netname)
			if err != nil {
				return
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

// NodeSet is the wwapi implmentation for updating node fields.
func NodeSet(set *wwapi.NodeSetParameter) (err error) {
	var nodeDB node.NodeYaml
	nodeDB, _, err = NodeSetParameterCheck(set, false)
	if err != nil {
		return
	}
	return nodeDbSave(&nodeDB)
}

// NodeSetParameterCheck does error checking on NodeSetParameter.
// Output to the console if console is true.
// TODO: Determine if the console switch does wwlog or not.
// - console may end up being textOutput?
func NodeSetParameterCheck(set * wwapi.NodeSetParameter, console bool) (nodeDB node.NodeYaml, nodeCount uint, err error) {
	//var err error
	//var count uint
	var setProfiles []string // TODO: Look at this. Is there an issue here?

	// TODO: Need these checks elsewhere.
	if set == nil {
		err = fmt.Errorf("Node set parameter is null")
		if console {
			fmt.Printf("%v\n", err)
			return
		}
	}

	if set.NodeNames == nil {
		err = fmt.Errorf("Node set parameter: NodeNames is null")
		if console {
			fmt.Printf("%v\n", err)
			return
		}
	}

	nodeDB, err = node.New()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not open node configuration: %s\n", err)
		return
	}

	nodes, err := nodeDB.FindAllNodes()
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not get node list: %s\n", err)
		return
	}

	// Note: This does not do expansion on the nodes.

	if set.AllNodes || (len(set.NodeNames) == 0 && len(nodes) > 0) {
		if console {
			fmt.Printf("\n*** WARNING: This command will modify all nodes! ***\n\n")
		}
	} else {
		nodes = node.FilterByName(nodes, set.NodeNames)
	}

	if len(nodes) == 0 {
		if console {
			fmt.Printf("No nodes found\n")
		}
		return
	}

	for _, n := range nodes {
		wwlog.Printf(wwlog.VERBOSE, "Evaluating node: %s\n", n.Id.Get())

		if set.Comment != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting comment to: %s\n", n.Id.Get(), set.Comment)
			n.Comment.Set(set.Comment)
		}

		if set.Container != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting container name to: %s\n", n.Id.Get(), set.Container)
			n.ContainerName.Set(set.Container)
		}

		if set.Init != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting init command to: %s\n", n.Id.Get(), set.Init)
			n.Init.Set(set.Init)
		}

		if set.Root != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting root to: %s\n", n.Id.Get(), set.Root)
			n.Root.Set(set.Root)
		}

		if set.AssetKey != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting asset key to: %s\n", n.Id.Get(), set.AssetKey)
			n.AssetKey.Set(set.AssetKey)
		}

		if set.Kernel != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting kernel to: %s\n", n.Id.Get(), set.Kernel)
			n.KernelVersion.Set(set.Kernel)
		}

		if set.KernelArgs != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting kernel args to: %s\n", n.Id.Get(), set.KernelArgs)
			n.KernelArgs.Set(set.KernelArgs)
		}

		if set.Cluster != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting cluster name to: %s\n", n.Id.Get(), set.Cluster)
			n.ClusterName.Set(set.Cluster)
		}

		if set.Ipxe != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting iPXE template to: %s\n", n.Id.Get(), set.Ipxe)
			n.Ipxe.Set(set.Ipxe)
		}

		if set.RuntimeOverlay != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting runtime overlay to: %s\n", n.Id.Get(), set.RuntimeOverlay)
			n.RuntimeOverlay.Set(set.RuntimeOverlay)
		}

		if set.SystemOverlay != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting system overlay to: %s\n", n.Id.Get(), set.SystemOverlay)
			n.SystemOverlay.Set(set.SystemOverlay)
		}

		if set.IpmiIpaddr != "" {
			newIpaddr := util.IncrementIPv4(set.IpmiIpaddr, nodeCount)
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI IP address to: %s\n", n.Id.Get(), newIpaddr)
			n.IpmiIpaddr.Set(newIpaddr)
		}

		if set.IpmiNetmask != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI netmask to: %s\n", n.Id.Get(), set.IpmiNetmask)
			n.IpmiNetmask.Set(set.IpmiNetmask)
		}

		if set.IpmiPort != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI port to: %s\n", n.Id.Get(), set.IpmiPort)
			n.IpmiPort.Set(set.IpmiPort)
		}

		if set.IpmiGateway != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI gateway to: %s\n", n.Id.Get(), set.IpmiGateway)
			n.IpmiGateway.Set(set.IpmiGateway)
		}

		if set.IpmiUsername != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI IP username to: %s\n", n.Id.Get(), set.IpmiUsername)
			n.IpmiUserName.Set(set.IpmiUsername)
		}

		if set.IpmiPassword != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI IP password to: %s\n", n.Id.Get(), set.IpmiPassword)
			n.IpmiPassword.Set(set.IpmiPassword)
		}

		if set.IpmiInterface != "" {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting IPMI IP interface to: %s\n", n.Id.Get(), set.IpmiInterface)
			n.IpmiInterface.Set(set.IpmiInterface)
		}

		if set.Discoverable {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting node to discoverable\n", n.Id.Get())
			n.Discoverable.SetB(true)
		}

		if set.Undiscoverable {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting node to undiscoverable\n", n.Id.Get())
			n.Discoverable.SetB(false)
		}

		if len(setProfiles) > 0 {
			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting profiles to: %s\n", n.Id.Get(), strings.Join(setProfiles, ","))
			n.Profiles = setProfiles
		}

		if len(set.ProfileAdd) > 0 {
			for _, p := range set.ProfileAdd {
				wwlog.Printf(wwlog.VERBOSE, "Node: %s, adding profile '%s'\n", n.Id.Get(), p)
				n.Profiles = util.SliceAddUniqueElement(n.Profiles, p)
			}
		}

		if len(set.ProfileDelete) > 0 {
			for _, p := range set.ProfileDelete {
				wwlog.Printf(wwlog.VERBOSE, "Node: %s, deleting profile '%s'\n", n.Id.Get(), p)
				n.Profiles = util.SliceRemoveElement(n.Profiles, p)
			}
		}

		if set.Netname != "" {
			if _, ok := n.NetDevs[set.Netname]; !ok {
				var nd node.NetDevEntry

				n.NetDevs[set.Netname] = &nd

				if set.Netdev == "" {
					n.NetDevs[set.Netname].Device.Set(set.Netname)
				}
			}
			var def bool = true

			// NOTE: This is overriding parameters passed in by the caller.
			set.Onboot = "yes"

			for _, n := range n.NetDevs {
				if n.Default.GetB() {
					def = false
				}
			}

			if def {
				set.NetDefault = "yes"
			}
		}

		if set.Netdev != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting net Device to: %s\n", n.Id.Get(), set.Netname, set.Netdev)
			n.NetDevs[set.Netname].Device.Set(set.Netdev)
		}

		if set.Ipaddr != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			newIpaddr := util.IncrementIPv4(set.Ipaddr, nodeCount)

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting Ipaddr to: %s\n", n.Id.Get(), set.Netname, newIpaddr)
			n.NetDevs[set.Netname].Ipaddr.Set(newIpaddr)
		}

		if set.Netmask != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting netmask to: %s\n", n.Id.Get(), set.Netname, set.Netmask)
			n.NetDevs[set.Netname].Netmask.Set(set.Netmask)
		}

		if set.Gateway != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting gateway to: %s\n", n.Id.Get(), set.Netname, set.Gateway)
			n.NetDevs[set.Netname].Gateway.Set(set.Gateway)
		}

		if set.Hwaddr != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting HW address to: %s\n", n.Id.Get(), set.Netname, set.Hwaddr)
			n.NetDevs[set.Netname].Hwaddr.Set(strings.ToLower(set.Hwaddr))
		}

		if set.Type != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting Type %s\n", n.Id.Get(), set.Netname, set.Type)
			n.NetDevs[set.Netname].Type.Set(set.Type)
		}

		if set.Onboot != "" {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			if set.Onboot == "yes" || set.Onboot == "y" || set.Onboot == "1" || set.Onboot == "true" {
				wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting ONBOOT\n", n.Id.Get(), set.Netname)
				n.NetDevs[set.Netname].OnBoot.SetB(true)
			} else {
				wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Unsetting ONBOOT\n", n.Id.Get(), set.Netname)
				n.NetDevs[set.Netname].OnBoot.SetB(false)
			}
		}

		if set.NetDefault != "" {
			if set.Netname == "" {
				err = fmt.Errorf("You must include the '--netname' option")
				wwlog.Printf(wwlog.ERROR, fmt.Sprintf("%v\n", err.Error()))
				return
			}

			if set.NetDefault == "yes" || set.NetDefault == "y" || set.NetDefault == "1" || set.NetDefault == "true" {

				// Set all other devices to non-default
				for _, n := range n.NetDevs {
					n.Default.SetB(false)
				}

				wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Setting DEFAULT\n", n.Id.Get(), set.Netname)
				n.NetDevs[set.Netname].Default.SetB(true)
			} else {
				wwlog.Printf(wwlog.VERBOSE, "Node: %s:%s, Unsetting DEFAULT\n", n.Id.Get(), set.Netname)
				n.NetDevs[set.Netname].Default.SetB(false)
			}
		}

		if set.NetdevDelete {
			err = checkNetNameRequired(set.Netname)
			if err != nil {
				return
			}

			if _, ok := n.NetDevs[set.Netname]; !ok {
				err = fmt.Errorf("Network device name doesn't exist: %s", set.Netname)
				wwlog.Printf(wwlog.ERROR, fmt.Sprintf("%v\n", err.Error()))
				return
			}

			wwlog.Printf(wwlog.VERBOSE, "Node: %s, Deleting network device: %s\n", n.Id.Get(), set.Netname)
			delete(n.NetDevs, set.Netname)
		}

		if len(set.Tags) > 0 {
			for _, t := range set.Tags {
				keyval := strings.SplitN(t, "=", 2)
				key := keyval[0]
				val := keyval[1]

				if _, ok := n.Tags[key]; !ok {
					var nd node.Entry
					n.Tags[key] = &nd
				}

				wwlog.Printf(wwlog.VERBOSE, "Node: %s, Setting Tag '%s'='%s'\n", n.Id.Get(), key, val)
				n.Tags[key].Set(val)
			}
		}
		if len(set.TagsDelete) > 0 {
			for _, t := range set.TagsDelete {
				keyval := strings.SplitN(t, "=", 1)
				key := keyval[0]

				if _, ok := n.Tags[key]; !ok {
					wwlog.Printf(wwlog.WARN, "Key does not exist: %s\n", key)
					os.Exit(1)
				}

				wwlog.Printf(wwlog.VERBOSE, "Node: %s, Deleting tag: %s\n", n.Id.Get(), key)
				delete(n.Tags, key)
			}
		}

		err := nodeDB.NodeUpdate(n)
		if err != nil {
			wwlog.Printf(wwlog.ERROR, "%s\n", err)
			os.Exit(1)
		}

		nodeCount++
	}
	return
}

// NodeSetPrompt prompt is a blocking confirmation prompt.
// Returns true on y or yes.
func NodeSetPrompt(label string) (yes bool) {

	prompt := promptui.Prompt{
		Label: label,
		IsConfirm: true,
	}

	result, _ := prompt.Run()
	if result == "y" || result == "yes" {
		yes = true
	}
	return
}

func checkNetNameRequired(netname string) (err error) {
	if netname == "" {
		err = fmt.Errorf("You must include the '--netname' option")
		wwlog.Printf(wwlog.ERROR, fmt.Sprintf("%v\n", err.Error()))
	}
	return
}

// nodeDbSave persists the nodeDB to disk and restarts warewulfd.
func nodeDbSave(nodeDB *node.NodeYaml) (err error) {
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