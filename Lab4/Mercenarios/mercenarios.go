package main

import (
	"context"
	"log"
	"time"

	pb "prueba1/proto"

	"google.golang.org/grpc"
)

type Mercenario struct {
	ID   string
	Name string
}

func NewMercenario(id, name string) *Mercenario {
	return &Mercenario{
		ID:   id,
		Name: name,
	}
}

func (m *Mercenario) InformarEstadoPreparacion(client pb.DirectorServiceClient) {
	req := &pb.PreparacionRequest{
		MercenarioId: m.ID,
		Nombre:       m.Name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.Preparacion(ctx, req)
	if err != nil {
		log.Fatalf("Error al informar el estado de preparación: %v", err)
	}
	log.Println("Estado de preparación informado correctamente")
}

func (m *Mercenario) TomarDecision(client pb.DirectorServiceClient, piso string) {
	req := &pb.DecisionRequest{
		MercenarioId: m.ID,
		Piso:         piso,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.Decision(ctx, req)
	if err != nil {
		log.Fatalf("Error al tomar una decisión: %v", err)
	}
	log.Println("Decisión tomada correctamente")
}

func (m *Mercenario) VerMontoDoshBank(client pb.DirectorServiceClient) {
	req := &pb.MontoRequest{
		Solicitud: "Solicitud de monto acumulado",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.ObtenerMonto(ctx, req)
	if err != nil {
		log.Fatalf("Error al obtener el monto del Dosh Bank: %v", err)
	}
	log.Printf("Monto acumulado en Dosh Bank: %d", res.Total)
}

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el Director: %v", err)
	}
	defer conn.Close()

	client := pb.NewDirectorServiceClient(conn)

	mercenario := NewMercenario("1", "Reverend David Alberts")

	mercenario.InformarEstadoPreparacion(client)
	mercenario.TomarDecision(client, "Piso_1")
	mercenario.VerMontoDoshBank(client)
}
