package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "Lab3/pb/helloworld"

	"google.golang.org/grpc"
)

// server es nuestra implementación del servicio HelloWorldServer
type server struct{}

// SayHello implementa el método SayHello del servicio HelloWorldServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hola " + in.Name}, nil
}

// main inicia un servidor gRPC y escucha las solicitudes

func main() {
	AT := 0
	MP := 0
	// Hacer lo de gRPC para recibir las consultas

	// Crear un servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()

	// Registrar el servidor para que el cliente pueda llamarlo
	pb.RegisterHelloWorldServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

	for true {
		// Esperar 5 segundos
		time.Sleep(50 * time.Millisecond)

		// Generar AT
		if AT < 40 {
			AT += 10
		} else {
			AT = 50
		}

		// Generar MP
		if MP < 15 {
			MP += 5
		} else {
			MP = 20
		}
	}
}

// Prints para agregar luego

// fmt.Print("Recepción de solicitud desde equipo", equipo, at_solicitado, "AT", mp_solicitado)
// fmt.Print(" -- APROBADA -- ")
// fmt.Print(" -- DENEGADA -- ")
// fmt.Println("AT EN SISTEMA:",AT,"; MP EN SISTEMA:", MP)
