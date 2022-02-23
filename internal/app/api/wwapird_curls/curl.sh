#! /usr/bin/env bash

# version
curl http://localhost:9871/version

# node list all
curl http://localhost:9871/v1/node

# node list one
curl http://localhost:9871/v1/node?nodeNames=testnode1 # this works! case sensitive

# node list some (See if this glob is working. Not clear that it is, but wwctl does work)
#curl 'http://localhost:9871/v1/node?NodeNames=testnode[1-2]'
# Works:
curl http://localhost:9871/v1/node?nodeNames=testnode%5B1-2%5D


# node add single discoverable node
curl -d '{"nodeNames": ["testApiNode0"], "discoverable": true}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node 

curl -d '{"nodeNames": ["testApiNode1"], "discoverable": true}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node 


# list the node we just added
curl http://localhost:9871/v1/node?nodeNames=testApiNode0

# node sets:
# This one does not seem to work in current main either.
curl -d '{"nodeNames": ["testApiNode0"], "undiscoverable": true}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 

# TODO: Try ipmiaddr. (Parameters are not being set.)
# wwctl does work though. Likely a parameter issue.
#curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220"}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node
#curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220", "updateMask": "ipmiIpAddr"}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 
#curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220", "updateMask": "*"}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 
#curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220", "updateMask": "\*"}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 
#curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220", "updateMask": [*]}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 
# below failed as well.
#curl -X PATCH http://localhost:9871/v1/node?nodeNames=testApiNode0&ipmiIpAddr=10.0.8.220 

# This gets me a little farther, but still no param data:
curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpAddr": "10.0.8.220", "updateMask": "ipmiIpAddr,nodeNames"}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node

#2022/02/22 20:54:13 request: *wwapiv1.NodeSetParameter, &wwapiv1.NodeSetParameter{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(0xc00010d040)}, sizeCache:0, unknownFields:[]uint8(nil), Comment:"", Container:"", Kernel:"", KernelArgs:"", Netname:"", Netdev:"", Ipaddr:"", Netmask:"", Gateway:"", Hwaddr:"", Type:"", Onboot:"", NetDefault:"", NetdevDelete:false, Cluster:"", Ipxe:"", InitOverlay:"", RuntimeOverlay:"", SystemOverlay:"", IpmiIpaddr:"", IpmiNetmask:"", IpmiPort:"", IpmiGateway:"", IpmiUsername:"", IpmiPassword:"", IpmiInterface:"", AllNodes:false, Profile:"", ProfileAdd:[]string(nil), ProfileDelete:[]string(nil), Force:false, Init:"", Discoverable:false, Undiscoverable:false, Root:"", Tags:[]string(nil), TagsDelete:[]string(nil), AssetKey:"", NodeNames:[]string(nil), UpdateMask:(*fieldmaskpb.FieldMask)(0xc0003bc1c0)}
#2022/02/22 20:54:13 *request.UpdateMask: *fieldmaskpb.FieldMask, &fieldmaskpb.FieldMask{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, sizeCache:0, unknownFields:[]uint8(nil), Paths:[]string{"ipmi_ip_addr", "node_names"}}

# Node set with post: WORKS!!!
curl -d '{"nodeNames": ["testApiNode0"], "ipmiIpaddr": "6.7.8.9"}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/nodeset 



# node delete single node
curl -X DELETE http://localhost:9871/v1/node?nodeNames=testApiNode0
curl -X DELETE http://localhost:9871/v1/node?nodeNames=testApiNode1
