#! /usr/bin/env bash

# version
#curl http://localhost:9871/version

# node list all
#curl http://localhost:9871/v1/node

# node list one
# This is listing all. Trace the parameter on the server.
# NodeNames is coming in as nil on the server.
#curl http://localhost:9871/v1/node?NodeNames=testnode1
#curl http://localhost:9871/v1/node?nodeNames=testnode1 # this works! case sensitive
curl http://localhost:9871/v1/node?nodeNames=testApiNode0 # this works! case sensitive

# node list some (See if this glob is working. Not clear that it is, but wwctl does work)
#curl 'http://localhost:9871/v1/node?NodeNames=testnode[1-2]'

# Works:
curl http://localhost:9871/v1/node?nodeNames=testnode%5B1-2%5D


# node add single node
#curl -d '{"nodeNames":"testApiNode0"}' -H "ContentType: application/json" -X POST http://localhost:9871/v1/node
#curl -d '{"nodeNames": "testApiNode0"}' -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt
#curl -d '{"nodeNames": "testApiNode0"}' -H "ContentType: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt
#curl -d '{"nodeNames": ["testApiNode0"]}' -H "ContentType: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt
#curl -d '{"NodeNames": ["testApiNode0"]}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt
#curl -d '{"NodeNames": {"NodeNames": ["testApiNode0"]}}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt

#curl -d '{"NodeNames": ["testApiNode0"], "Discoverable": true}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt


#curl -d '{"nodeNames": ["testApiNode0"], "Discoverable": true}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt
curl -d '{"nodeNames": ["testApiNode0"], "discoverable": true}' -H "Content-Type: application/json" -X POST http://localhost:9871/v1/node --trace-ascii tmp.txt


# node delete single node
#curl -X DELETE http://localhost:9871/v1/node?nodeNames=testnode1
#curl -X DELETE --trace-ascii tmp.txt http://localhost:9871/v1/node?force=true&nodeNames=testnode1
curl -X DELETE --trace-ascii tmp.txt http://localhost:9871/v1/node?nodeNames=testApiNode0
