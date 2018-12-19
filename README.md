# Goals
1. Write 2 programs using "Go" language (golang.org)
    - one program listens to a network endpoint
    - the other sends some data there.
    - Data serialization format/mechanism is up to you.
3. create a docker image for each program.
4. test the 2 docker containers so that they talk to each other.
5. test with 2-4 docker containers of the sender.

## Usage
build the docker containers with `make -B`

run with `./run.sh`


## Data format ##
Each sender will include an id and a randomly generated number, formatted as json.

The sender will send out an event randomly every <2 seconds

The receiver will log those along with a timestamp of when the log is received and append to a file.
