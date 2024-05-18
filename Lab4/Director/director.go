package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	pb "prueba1/proto"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var conn *amqp.Connection
var ch *amqp.Channel
var serviceClient pb.BankServiceClient
var grpcInitialized bool
var grpcMutex sync.Mutex

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
	grpcMutex.Lock()
	defer grpcMutex.Unlock()

	if !grpcInitialized {
		log.Println("gRPC connection not initialized")
		return 0
	}

	res, err := serviceClient.GetTotal(context.Background(), &pb.GetTotalRequest{})
	if err != nil {
		log.Fatalf("Failed to get total: %s", err)
	}
	return res.Total
}

func initializeGrpcConnection() {
	conn, err := grpc.Dial("dist043.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect with server: %s", err)
	}

	serviceClient = pb.NewBankServiceClient(conn)
	grpcInitialized = true
	log.Println("gRPC connection initialized")
}

func main() {
	initializeRabbitMQConnection()

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

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Initialize gRPC connection on first message
			if !grpcInitialized {
				initializeGrpcConnection()

				// Perform initial gRPC request
				totalAmount := getTotalAmount()
				fmt.Printf("Monto acumulado en Dosh Bank: %d\n", totalAmount)
			}

			// Send death message
			body := string(d.Body)
			components := strings.Split(body, ",")
			mercenary := components[0]
			floor := components[1]

			sendDeathMessage(mercenary, floor)
		}
	}()

	// Simulate sending messages
	sendDeathMessage("Mercenario1", "Piso_1")
	sendDeathMessage("Mercenario2", "Piso_2")
	sendDeathMessage("Mercenario3", "Piso_3")

}
