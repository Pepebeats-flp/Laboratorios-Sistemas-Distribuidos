package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func createFile() {
	file, err := os.Create("dosh_bank.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
}

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://dist:dist@dist041.inf.santiago.usm.cl:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"dosh_bank2", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

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

	// Loop to receive messages
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		body := string(d.Body)

		// Split the message into its components
		components := strings.Split(body, ",")
		mercenary := components[0]
		floor := components[1]

		// Initialize the amount to 0 each time
		amount := 0

		// Write the components to the file
		writeToFile(mercenary, floor, fmt.Sprintf("%d", amount))
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
