package main

import (
	"Lab3/servidor_central"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Naves struct {
	ID int
	AT int
	MP int
}

func main() {
	// Conexión con el servidor
	puerto := ":6969"
	conn, err := grpc.Dial(puerto, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error al conectar en el puerto:"+puerto+" %v", err)
	}
	defer conn.Close()

	cliente := servidor_central.NewServidorCentralClient(conn)

	// Crear los 4 clientes
	naves := make([]Naves, 4)
	for i := 1; i <= 4; i++ {
		at := rand.Intn(11) + 20
		mp := rand.Intn(6) + 10

		naves[i-1] = Naves{ID: i, AT: at, MP: mp}
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {

		go func(i int) {
			defer wg.Done()
			time.Sleep(10 * time.Second)
			peticion := servidor_central.PedirMunicion{
				Id: int32(naves[i].ID),
				At: int32(naves[i].AT),
				Mp: int32(naves[i].MP),
			}
			for {
				respuesta, err := cliente.SolicitudMunicion(context.Background(), &peticion)
				if err != nil {
					log.Fatalf("No se pudo realizar la solicitud: %v", err)
				}

				auxiliar := respuesta.Respuesta
				fmt.Print("Solicitando AT: ", naves[i].AT, " y MP: ", naves[i].MP)
				if auxiliar == 1 {
					fmt.Printf("Resolución: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicación\n")
					return
				} else {
					fmt.Printf("Resolución: -- DENEGADA -- ; Reintentando en 3 segs...\n")
					time.Sleep(3 * time.Second)
				}
			}
		}(i)
	}
	wg.Wait()
}
