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
	Live bool
}

// NewMercenario crea una nueva instancia de Mercenario
func NewMercenario(id, name string) *Mercenario {
	return &Mercenario{
		ID:   id,
		Name: name,
		Live: true,
	}
}

// InformarEstadoPreparacion envía una solicitud al Director para informar el estado de preparación
func (m *Mercenario) InformarEstadoPreparacion(client pb.DirectorServiceClient) {
	req := &pb.PreparacionRequest{
		MercenarioId: m.ID,
		Nombre:       m.Name,
		Preparado:    true,
	}
	resp, err := client.Preparacion(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al informar preparación: %v", err)
	}
	if resp.Mensaje == "" {
		log.Printf("Preparación nula")
	}
}

// TomarDecision toma una decisión en base a la situación del piso actual
func (m *Mercenario) TomarDecision(client pb.DirectorServiceClient) {
	req := &pb.DecisionRequest{
		MercenarioId: m.Name,   // Nombre del mercenario
		Piso:         "Piso_1", // Ejemplo de piso CAMBIAR

	}
	resp, err := client.Decision(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al tomar decisión: %v", err)
	}
	if resp.Mensaje == "" {
		log.Printf("Decisión nula")
	}

}

// VerMontoDoshBank solicita al Director el monto acumulado en el Dosh Bank
func VerMontoDoshBank(client pb.DirectorServiceClient) {
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

// Enviar mensaje al Director para el Piso 1
func EnviarMensajePiso1(client pb.DirectorServiceClient, mercenarios []*pb.Mercenario) {
	req := &pb.Piso1Request{
		Mercenarios: mercenarios,
	}
	// Enviar el mensaje al Director
	// Manejar la respuesta si es necesario

}

// Enviar mensaje al Director para el Piso 2
func EnviarMensajePiso2(client pb.DirectorServiceClient, decisiones []*pb.Decision) {
	req := &pb.Piso2Request{
		Decisiones: decisiones,
	}
	// Enviar el mensaje al Director
	// Manejar la respuesta si es necesario
}

// Enviar mensaje al Director para el Piso 3
func EnviarMensajePiso3(client pb.DirectorServiceClient, numeros []int32) {
	req := &pb.Piso3Request{
		Numeros: numeros,
	}
	// Enviar el mensaje al Director
	// Manejar la respuesta si es necesario
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
		log.Printf("Mercenario %s (%s) informa que está preparado", mercenario.Name, mercenario.ID)
	}

	// Simular un retraso antes de eliminar a los mercenarios
	time.Sleep(2 * time.Second)

	// Simular la eliminación de los mercenarios
	for _, mercenario := range mercenarios {
		mercenario.TomarDecision(directorClient)
		log.Printf("Mercenario %s toma una decisión en el piso %s", mercenario.Name, "Piso_1")
	}

	// Simular un retraso antes de pedir el monto
	time.Sleep(2 * time.Second)

	// Solicitar ver el monto acumulado en el Dosh Bank una vez
	VerMontoDoshBank(directorClient)

}
