package main

import (
	pb "Lab4/Dosh_Bank"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// Create a file to store the messages
func createFile() {
	file, err := os.Create("dosh_bank.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
}

// Write the components to the file
func writeToFile(mercenary string, floor string, amount string) {
	file, err := os.OpenFile("dosh_bank.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	_, err = file.WriteString("- " + mercenary + " " + floor + " " + amount + "\n")
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

// Read the file
func readFromFile() string {
	file, err := os.Open("dosh_bank.txt")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Cannot get file info", err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	return string(data)
}

// Fail on error
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// DoshBankServer implementa el servidor gRPC para manejar solicitudes del monto acumulado
type DoshBankServer struct{}

// GetTotalAmount devuelve el monto acumulado actual
func (s *DoshBankServer) GetTotalAmount(ctx context.Context, request *pb.TotalAmountRequest) (*pb.TotalAmountResponse, error) {
	// Aquí puedes implementar la lógica para obtener el monto acumulado del archivo o de cualquier otra fuente

	// Leer el archivo y obtener el ultimo monto
	fileContent := readFromFile()
	lines := strings.Split(fileContent, "\n")
	lastLine := lines[len(lines)-2]
	components := strings.Split(lastLine, " ")
	totalAmount, err := strconv.Atoi(components[2])
	if err != nil {
		log.Fatalf("Cannot convert amount to integer", err)
	}
	// Devolver el monto acumulado
	return &pb.TotalAmountResponse{Amount: totalAmount}, nil
}

func main() {

	// Iniciar el servidor gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDoshBankServer(grpcServer, &DoshBankServer{})
	log.Println("gRPC server started on port 50051")
	grpcServer.Serve(listener)

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://dist:dist@dist041.inf.santiago.usm.cl:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// Close the connection at the end
	defer conn.Close()
	// Open a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// Close the channel at the end
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"eliminated_mercenaries", // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Create file at the beginning
	createFile()

	// initial amount
	amount := 0

	// Print initial amount
	log.Printf(" [x] Amount: %d\n", amount)

	// Loop to receive messages
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		body := string(d.Body)

		// Split the message into its components
		components := strings.Split(body, ",")
		mercenary := components[0]
		floor := components[1]

		// Increment the amount
		amount += 100000000

		// Print the updated amount
		log.Printf(" [x] Amount: %d\n", amount)

		// Write the components to the file
		writeToFile(mercenary, floor, fmt.Sprintf("%d", amount))
	}

	// Read the file
	log.Printf(" [x] File content: %s", readFromFile())
	// reinitialize the amount
	amount = 0
}
