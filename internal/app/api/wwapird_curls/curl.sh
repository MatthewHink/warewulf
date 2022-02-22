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

# list the node we just added
curl http://localhost:9871/v1/node?nodeNames=testApiNode0

# node sets:
curl -d '{"nodeNames": ["testApiNode0"], "discoverable": false}' -H "Content-Type: application/json" -X PATCH http://localhost:9871/v1/node 

# node delete single node
curl -X DELETE http://localhost:9871/v1/node?nodeNames=testApiNode0
