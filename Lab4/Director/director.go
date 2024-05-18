package main

import (
	"context"
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

func sendDeathMessage(mercenaryName string, floor string) {
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

	body := mercenaryName + "," + floor

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

	log.Printf("Mercenario %s muerto en el piso %s", mercenaryName, floor[5:])
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
	log.Printf("Mercenario %s está listo", req.Nombre)
	return &pb.PreparacionResponse{Mensaje: "Preparación recibida correctamente"}, nil
}

// Implementación del servicio Decision
func (s *DirectorServer) Decision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	sendDeathMessage(req.Nombre, req.Piso)
	return &pb.DecisionResponse{Mensaje: "Decisión recibida correctamente"}, nil
}

// Implementación del servicio ObtenerMonto
func (s *DirectorServer) ObtenerMonto(ctx context.Context, req *pb.MontoRequest) (*pb.MontoResponse, error) {
	log.Println("Solicitud de monto recibida")
	total := getTotalAmount()
	log.Printf("Monto acumulado en Dosh Bank: %d", total)
	return &pb.MontoResponse{Total: total}, nil
}

func main() {
	// Inicializar conexión con RabbitMQ
	initializeRabbitMQConnection()

	// Conectar al servidor gRPC del banco
	bankConn, err := grpc.Dial("dist097.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with bank server " + err.Error())
	}
	defer bankConn.Close()

	serviceClient = pb.NewBankServiceClient(bankConn)

	// Iniciar el servidor gRPC para el Director-mercenarios
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServiceServer(s, &DirectorServer{})
	log.Printf("Director server started at port :50052")

	// Mantener el servidor gRPC en ejecución
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
