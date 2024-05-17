// director_client.go
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/Pepebeats-flp/Laboratorios-Sistemas-Distribuidos/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

var conn *amqp.Connection
var ch *amqp.Channel

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

func main() {
	// Conectar al servidor gRPC del banco
	conn, err := grpc.Dial("dist043.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	bankClient := pb.NewBankServiceClient(conn)

	// Solicitar el monto acumulado al banco
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := bankClient.GetTotal(ctx, &pb.GetTotalRequest{})
	if err != nil {
		log.Fatalf("Failed to get total: %v", err)
	}

	fmt.Printf("Total amount in bank: %d\n", res.Total)

	// Inicializar conexión a RabbitMQ
	initializeRabbitMQConnection()

	sendDeathMessage("Mercenario1", "Piso_1")
	sendDeathMessage("Mercenario2", "Piso_2")
	sendDeathMessage("Mercenario3", "Piso_3")

	// Esperar para evitar que el programa termine inmediatamente
	select {}
}
