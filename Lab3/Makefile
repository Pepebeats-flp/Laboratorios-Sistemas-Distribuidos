.PHONY: client proto docker server

proto:
	protoc -I=. --go_out=. --go-grpc_out=. servidor_central/protocolo.proto

docker: proto
	sudo docker build -t lab3 .

server: docker
	sudo docker run -p 6969:6969 lab3

run:
	sudo docker run -p 6969:6969 lab3

client:
	go run client.go
