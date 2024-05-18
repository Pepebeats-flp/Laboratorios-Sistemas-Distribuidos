package main

import (
	"context"
	"log"
	"time"

	pb "prueba1/proto"

	"google.golang.org/grpc"
)

// Mercenario representa a un participante en la misión
type Mercenario struct {
	ID   string
	Name string
	// Otros campos necesarios
}

// NewMercenario crea una nueva instancia de Mercenario
func NewMercenario(id, name string) *Mercenario {
	return &Mercenario{
		ID:   id,
		Name: name,
		// Inicializar otros campos si es necesario
	}
}

// InformarEstadoPreparacion envía una solicitud al Director para informar el estado de preparación
func (m *Mercenario) InformarEstadoPreparacion(client pb.DirectorServiceClient) {
	req := &pb.PreparacionRequest{
		MercenarioId: m.ID,
		Nombre:       m.Name,
	}
	resp, err := client.Preparacion(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al informar preparación: %v", err)
	}
	log.Printf("Respuesta del Director: %s", resp.Mensaje)
}

// TomarDecision toma una decisión en base a la situación del piso actual
func (m *Mercenario) TomarDecision(client pb.DirectorServiceClient) {
	req := &pb.DecisionRequest{
		MercenarioId: m.ID,
		Piso:         "Piso_1", // Ejemplo de piso
	}
	resp, err := client.Decision(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al tomar decisión: %v", err)
	}
	log.Printf("Respuesta del Director: %s", resp.Mensaje)
}

// VerMontoDoshBank solicita al Director el monto acumulado en el Dosh Bank
func (m *Mercenario) VerMontoDoshBank(client pb.DirectorServiceClient) {
	req := &pb.MontoRequest{
		Solicitud: "VerMonto",
	}
	resp, err := client.ObtenerMonto(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al ver monto: %v", err)
	}
	log.Printf("Monto acumulado en Dosh Bank: %d", resp.Total)
}

// Función de inicialización para establecer una conexión gRPC con el Director
func initializeDirectorConnection() pb.DirectorServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el Director: %v", err)
	}
	client := pb.NewDirectorServiceClient(conn)
	return client
}

func main() {
	// Inicializar la conexión con el Director
	directorClient := initializeDirectorConnection()

	// Crear instancias de Mercenario
	mercenarios := []*Mercenario{
		NewMercenario("1", "Reverend David Alberts"),
		NewMercenario("2", "Police Constable Rob Briar"),
		NewMercenario("3", "Rae Higgins"),
		NewMercenario("4", "D.A.R."),
		NewMercenario("5", "Mr. Foster"),
		NewMercenario("6", "Donovan Neal"),
		NewMercenario("7", "Lieutenant William 'Bill' Masterson"),
		NewMercenario("8", "DJ Scully"),
	}

	// Informar estado de preparación para cada mercenario
	for _, mercenario := range mercenarios {
		mercenario.InformarEstadoPreparacion(directorClient)
	}

	// Simular un retraso antes de eliminar a los mercenarios
	time.Sleep(2 * time.Second)

	// Simular la eliminación de los mercenarios
	for _, mercenario := range mercenarios {
		mercenario.TomarDecision(directorClient)
	}

	// Simular un retraso antes de pedir el monto
	time.Sleep(2 * time.Second)

	// Solicitar ver el monto acumulado en el Dosh Bank
	for _, mercenario := range mercenarios {
		mercenario.VerMontoDoshBank(directorClient)
	}
}
