#!/bin/bash

# clear log and ensure it exists
mkdir -p log
echo > log/event.log

# expose the log file from the receiver container
docker run --network host -v $(pwd)/log/:/app/log -t receiver:latest &

docker run --network host -t sender:latest -id=1 &
docker run --network host -t sender:latest -id=2 &
docker run --network host -t sender:latest -id=3 &

tail -f log/event.log
