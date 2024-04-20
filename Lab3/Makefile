.PHONY: all server client clean

all: server client

server:
	cd servidor_central && go build -o ../server server.go
	./server &

client:
	go build -o client client.go
	./client

clean:
	rm -f server client