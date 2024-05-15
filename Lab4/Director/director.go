package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Funcion q informa muerte de un mercenario
// Se conecta a la cola de mensajes y envia un mensaje con el nombre del mercenario y el piso en el que murio
// El mensaje es enviado a la cola "dosh_bank2"
func sendDeathMessage(mercenary string, floor string) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := mercenary + "," + floor

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

	log.Printf(" [x] Mercenary %s died on floor %s", mercenary, floor)

}
func main() {
	sendDeathMessage("Mercenario1", "Piso_1")
	sendDeathMessage("Mercenario2", "Piso_2")
	sendDeathMessage("Mercenario3", "Piso_3")
}
