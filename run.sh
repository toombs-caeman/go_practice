#!/bin/bash

# clear log and ensure it exists
mkdir -p log
echo > log/event.log

# get the ip address of the docker bridge
HOSTIP="$(ifconfig docker0| awk '/inet addr/{print substr($2,6)}'):8080/"

# expose the log file from the receiver container
docker run -p 8080:80 -v $(pwd)/log/:/app/log -t receiver:latest & 

docker run -p 8081:80 -t sender:latest -id=1 -url="$HOSTIP" & 
#docker run -p 8082:80 -t sender:latest -id=2 -url=$HOSTIP & 

tail -f log/event.log
