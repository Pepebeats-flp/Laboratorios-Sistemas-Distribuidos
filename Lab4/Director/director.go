package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var ch *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func initializeRabbitMQConnection() {
	// Connect to RabbitMQ if not already connected
	if conn == nil {
		var err error
		conn, err = amqp.Dial("amqp://dist:dist@dist041.inf.santiago.usm.cl:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	// Open a channel if not already opened
	if ch == nil {
		var err error
		ch, err = conn.Channel()
		failOnError(err, "Failed to open a channel")
	}
}

func sendDeathMessage(mercenary string, floor string) {
	// Declare a queue if not already declared
	q, err := ch.QueueDeclare(
		"dosh_bank2", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := mercenary + "," + floor

	// Publish a message
	err = ch.PublishWithContext(ctx,
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

	initializeRabbitMQConnection()

	sendDeathMessage("Mercenario1", "Piso_1")
	sendDeathMessage("Mercenario2", "Piso_2")
	sendDeathMessage("Mercenario3", "Piso_3")

	// Close the channel and connection when no longer needed
	defer ch.Close()
	defer conn.Close()
}
