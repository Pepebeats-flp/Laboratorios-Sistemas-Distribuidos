// 5.3 Dosh Bank
// Es el encargado de mantener el conteo del monto acumulado por los mercenarios en la misi ́on.
// Debe crear un archivo txt donde se registre cada uno de los mercenarios eliminados y el monto acumulado actual.

//El Dosh Bank debe ser capaz de:
// • Registrar cada uno de los mercenarios eliminados en el archivo txt de la siguiente forma: – Mercenario Numero piso Monto acumulado actual
//– D.A.R. Piso 1 100000000
//– Mr.Foster Piso 2 200000000
//• Responder a las peticiones sobre el monto actual acumulado
//Este proceso debe estar corriendo solamente en una de las m ́aquinas virtuales.
// Debe proce- sar de manera as ́ıncrona, mediante RabbitMQ, el registro de mercenarios eliminados,
// pero de manera s ́ıncrona responder a la petici ́on del monto acumulado.

package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
