all: sender receiver

sender: sender/main.go
	cd sender; docker build -t sender .

receiver: receiver/main.go
	cd receiver; docker build -t receiver .
