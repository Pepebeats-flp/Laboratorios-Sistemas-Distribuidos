package main

import (
	"log"

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
	// Código para enviar la solicitud gRPC al Director
	// Usar 'client' para llamar a los métodos definidos en el servicio Director
}

// TomarDecision toma una decisión en base a la situación del piso actual
func (m *Mercenario) TomarDecision(client pb.DirectorServiceClient) {
	// Código para tomar una decisión y enviarla al Director
}

// VerMontoDoshBank solicita al Director el monto acumulado en el Dosh Bank
func (m *Mercenario) VerMontoDoshBank(client pb.DirectorServiceClient) {
	// Código para solicitar al Director el monto acumulado en el Dosh Bank
}

// Función de inicialización para establecer una conexión gRPC con el Director
func initializeDirectorConnection() pb.DirectorServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el Director: %v", err)
	}
	return pb.NewDirectorServiceClient(conn)
}

func main() {
	// Crear un canal para la señalización de finalización
	done := make(chan struct{})

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

	// Tomar decisiones en cada piso
	// (Esta parte se implementará más tarde después de definir la lógica de los pisos)
	for _, mercenario := range mercenarios {
		mercenario.TomarDecision(directorClient)
	}

	// Solicitar ver el monto acumulado en el Dosh Bank
	for _, mercenario := range mercenarios {
		mercenario.VerMontoDoshBank(directorClient)
	}

	// Esperar a que se reciba una señal para finalizar
	<-done
}
