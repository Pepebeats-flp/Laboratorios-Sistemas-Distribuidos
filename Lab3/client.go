package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

func random_integer(n int, m int) int {
	return n + rand.Intn(m-n)
}

func main() {

	// Conexión con el servidor
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer conn.Close()

	// Crear un cliente
	c := pb.NewHelloWorldClient(conn)

	// Llamada al servidor
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "usuario"})
	if err != nil {
		log.Fatalf("Error al llamar al servidor: %v", err)
	}
	log.Printf("Respuesta del servidor: %s", r.Message)

	start := time.Now().Second()
	for time.Now().Second()-start < 10 {
		// Esperar 7 segundos
		time.Sleep(70 * time.Millisecond)
	}

	AT := random_integer(20, 30)
	MP := random_integer(10, 15)
	acc := false
	for !acc {
		// Esperar 3 segundos
		time.Sleep(70 * time.Millisecond)
		acc = true // Consulta al servidor (borrar true al agregar consulta)

		fmt.Print("Solicitando", AT, "y", MP, "MP ; ")
		if acc {
			fmt.Println("Resolución: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicación")
		} else {
			fmt.Println("Resolución: -- DENEGADA -- ; Reintentando en 3 segs...")
		}

	}

}
