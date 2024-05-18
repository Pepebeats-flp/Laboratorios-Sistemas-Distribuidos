package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	log.Println("Mercenario eliminado: ", mercenaryName+" en el piso "+floor[5:])
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

// Piso1 maneja la lógica del piso 1
func (s *DirectorServer) Piso1(ctx context.Context, req *pb.Piso1Request) (*pb.Piso1Response, error) {
	log.Println("Iniciando Piso 1: Entrada al infierno")

	// Manejar la lógica del piso 1 aquí
	// Por ejemplo, calcular probabilidades de supervivencia para cada mercenario

	// Simulación de cálculo de probabilidades
	probabilidades := make(map[string]int32)
	for _, mercenario := range req.Mercenarios {
		probEscopeta := rand.Int31n(100)
		probRifle := rand.Int31n(100 - probEscopeta)
		probPunos := 100 - probEscopeta - probRifle
		probabilidades[mercenario.Id] = probEscopeta
		log.Println("Probpunos: ", probPunos)
		// Enviar mensaje al mercenario con la probabilidad de supervivencia
	}

	// Simulación de eliminación de mercenarios basada en probabilidades
	// Eliminar mercenarios basados en las probabilidades calculadas
	// Aquí puedes enviar mensajes de eliminación a los mercenarios que hayan fallado
	// Simulación de eliminación de mercenarios
	log.Println("Finalizando Piso 1")
	return &pb.Piso1Response{}, nil
}

// Piso2 maneja la lógica del piso 2
func (s *DirectorServer) Piso2(ctx context.Context, req *pb.Piso2Request) (*pb.Piso2Response, error) {
	log.Println("Iniciando Piso 2: Trampas y Traiciones")

	// Manejar la lógica del piso 2 aquí
	// Por ejemplo, calcular la salida correcta al azar

	// Simulación de cálculo de salida correcta al azar
	salidaCorrecta := rand.Intn(2)
	log.Printf("La salida correcta es el camino %c", 'A'+salidaCorrecta)

	// Simulación de eliminación de mercenarios basada en salida correcta
	// Aquí puedes enviar mensajes de eliminación a los mercenarios que hayan fallado

	log.Println("Finalizando Piso 2")
	return &pb.Piso2Response{}, nil
}

// Piso3 maneja la lógica del piso 3
func (s *DirectorServer) Piso3(ctx context.Context, req *pb.Piso3Request) (*pb.Piso3Response, error) {
	log.Println("Iniciando Piso 3: Confrontación Final")

	// Simulación de generación de número aleatorio del Patriarca
	patriarca := rand.Intn(15) + 1
	log.Printf("Número del Patriarca: %d", patriarca)

	// Simulación de rondas y aciertos
	aciertos := 0
	for i := 0; i < 5; i++ {
		numeroMercenario := rand.Intn(15) + 1
		log.Printf("Número elegido por el mercenario: %d", numeroMercenario)
		if numeroMercenario == patriarca {
			aciertos++
		}
	}

	// Verificar el resultado final
	if aciertos >= 2 {
		log.Println("Los mercenarios han salido victoriosos del enfrentamiento con el Patriarca!")
	} else {
		log.Println("Los mercenarios han sido eliminados en el enfrentamiento con el Patriarca.")
	}

	log.Println("Finalizando Piso 3")
	return &pb.Piso3Response{}, nil
}

// Implementación del servicio Preparacion
func (s *DirectorServer) Preparacion(ctx context.Context, req *pb.PreparacionRequest) (*pb.PreparacionResponse, error) {
	log.Println("Mercenario %s (%s) informa que está preparado", req.Nombre, req.MercenarioId)
	return &pb.PreparacionResponse{Mensaje: "Preparación recibida correctamente"}, nil
}

// Implementación del servicio Decision
func (s *DirectorServer) Decision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	log.Println("Mercenario %s toma una decisión en el piso %s", req.MercenarioId, req.Piso)
	sendDeathMessage(req.MercenarioId, req.Piso)
	return &pb.DecisionResponse{Mensaje: "Decisión recibida correctamente"}, nil
}

// Implementación del servicio ObtenerMonto
func (s *DirectorServer) ObtenerMonto(ctx context.Context, req *pb.MontoRequest) (*pb.MontoResponse, error) {
	log.Println("Solicitud de monto acumulado en Dosh Bank")
	total := getTotalAmount()
	fmt.Printf("Monto acumulado en Dosh Bank: %d\n", total)
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
