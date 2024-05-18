package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "prueba1/proto"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var conn *amqp.Connection
var ch *amqp.Channel
var serviceClient pb.BankServiceClient

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func initializeRabbitMQConnection() {
	var err error
	conn, err = amqp.Dial("amqp://dist:dist@dist043.inf.santiago.usm.cl:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func sendDeathMessage(mercenary string, floor string) {
	// Declare a queue if not already declared
	q, err := ch.QueueDeclare(
		"eliminated_mercenaries", // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := mercenary + "," + floor

	// Publish a message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Mercenary %s died on floor %s", mercenary, floor[5:])
}

func getTotalAmount() int32 {
	res, err := serviceClient.GetTotal(context.Background(), &pb.GetTotalRequest{})
	if err != nil {
		log.Fatalf("Failed to get total: %s", err)
	}
	return res.Total
}

type DirectorServer struct {
	pb.UnimplementedDirectorServiceServer
}

// InformarPreparacion implementa el método del servicio Director para recibir información de preparación de los mercenarios
func (s *DirectorServer) InformarPreparacion(ctx context.Context, req *pb.PreparacionRequest) (*pb.PreparacionResponse, error) {
	log.Printf("Se recibió un mensaje de preparación del mercenario %s\n", req.MercenarioId)
	return &pb.PreparacionResponse{}, nil
}

// TomarDecision implementa el método del servicio Director para recibir las decisiones de los mercenarios
func (s *DirectorServer) TomarDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	log.Printf("Se recibió una decisión del mercenario %s: %s\n", req.MercenarioId, req.Decision)
	return &pb.DecisionResponse{}, nil
}

// SolicitarMonto implementa el método del servicio Director para recibir solicitudes de ver el monto del Dosh Bank
func (s *DirectorServer) SolicitarMonto(ctx context.Context, req *pb.MontoRequest) (*pb.MontoResponse, error) {
	log.Printf("Se recibió una solicitud de ver el monto del Dosh Bank del mercenario %s\n", req.MercenarioId)
	// Aquí puedes realizar la lógica para obtener el monto del Dosh Bank y enviarlo como respuesta
	return &pb.MontoResponse{Monto: 1000}, nil // Ejemplo: devolver un monto de 1000
}

func main() {
	// Iniciar el servidor gRPC del Director
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServiceServer(s, &DirectorServer{})
	log.Printf("Servidor del Director iniciado en el puerto :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

	initializeRabbitMQConnection()

	sendDeathMessage("Mercenario1", "Piso_1")
	sendDeathMessage("Mercenario2", "Piso_2")
	sendDeathMessage("Mercenario3", "Piso_3")

	// Conectar al servidor gRPC del banco
	conn, err := grpc.Dial("dist043.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	defer conn.Close()

	serviceClient = pb.NewBankServiceClient(conn)

	// Realizar consulta inicial del monto acumulado
	totalAmount := getTotalAmount()
	fmt.Printf("Monto acumulado en Dosh Bank: %d\n", totalAmount)

	// Llamar a getTotalAmount en cualquier momento
	totalAmount = getTotalAmount()
	fmt.Printf("Monto acumulado en Dosh Bank después de enviar mensajes: %d\n", totalAmount)

}
