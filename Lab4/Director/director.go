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
	conn, err = amqp.Dial("amqp://dist:dist@dist097.inf.santiago.usm.cl:5672/")
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

// Implementación del servicio Preparacion
func (s *DirectorServer) Preparacion(ctx context.Context, req *pb.PreparacionRequest) (*pb.PreparacionResponse, error) {
	log.Printf("Solicitud de preparación recibida para el mercenario ID: %s, Nombre: %s", req.MercenarioId, req.Nombre)
	return &pb.PreparacionResponse{Mensaje: "Preparación recibida correctamente"}, nil
}

// Implementación del servicio Decision
func (s *DirectorServer) Decision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	log.Printf("Solicitud de decisión recibida para el mercenario ID: %s, Piso: %s", req.MercenarioId, req.Piso)
	return &pb.DecisionResponse{Mensaje: "Decisión recibida correctamente"}, nil
}

// Implementación del servicio ObtenerMonto
func (s *DirectorServer) ObtenerMonto(ctx context.Context, req *pb.MontoRequest) (*pb.MontoResponse, error) {
	// Conectar al servidor gRPC del banco
	conn, err := grpc.Dial("dist097.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	defer conn.Close()

	serviceClient = pb.NewBankServiceClient(conn)
	log.Printf("Solicitud de monto recibida: %s", req.Solicitud)
	// Aquí debes implementar la lógica para obtener el monto del Dosh Bank
	total := getTotalAmount()
	fmt.Printf("Monto acumulado en Dosh Bank: %d\n", total)
	return &pb.MontoResponse{Total: total}, nil
}

func main() {
	// Iniciar el servidor gRPC para el Director-mercenarios
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServiceServer(s, &DirectorServer{})
	log.Printf("Director server started at port :50052")

	initializeRabbitMQConnection()

	// Mantener el servidor gRPC en ejecución
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
