package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://dist:dist@10.35.169.51:5672/")

	if err != nil {
		log.Printf(" [x] Failed to connect to RabbitMQ")
	} else {
		log.Printf(" [x] Connected to RabbitMQ")
	}

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	if err != nil {
		log.Printf(" [x] Failed to open a channel")
	} else {
		log.Printf(" [x] Opened a channel")
	}

	q, err := ch.QueueDeclare(
		"dosh_bank", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if err != nil {
		log.Printf(" [x] Failed to declare a queue")
	} else {
		log.Printf(" [x] Declared a queue")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err != nil {
		log.Printf(" [x] Failed to create a context")
	} else {
		log.Printf(" [x] Created a context")
	}

	body := "Mercenario_1,Piso_1"

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

	if err != nil {
		log.Printf(" [x] Sent %s", body)
	} else {
		log.Printf(" [x] Failed to send %s", body)
	}

}
