package main

import (
	"Lab3/servidor_central"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Hacer lo de gRPC para recibir las consultas
	port := ":6969"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al escuchar el puerto:"+port+" %v", err)
	}

	s := servidor_central.Server{}
	s.AT = 0
	s.MP = 0

	// crear municion
	go func() {
		for {
			time.Sleep(5 * time.Second)
			s.ActualizarMunicion(10, 5)
		}
	}()

	grpcServer := grpc.NewServer()

	servidor_central.RegisterServidorCentralServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al ejecutar grcp en el puerto:"+port+" %v", err)
	}
}
