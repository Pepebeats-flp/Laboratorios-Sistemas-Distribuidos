package main

import (
	"fmt"
	"math/rand"
	"net"
)

func random_integer(n int) int {
	return rand.Intn(n)
}

// Funcion que retorne un planeta al azar
func planet_asignation(n int) string {
	// Definir las claves "PA", "PB", ..., "PF"
	claves := []string{"PA", "PB", "PC", "PD", "PE", "PF"}
	// Elegir un planeta al azar y retornarlo
	planeta := claves[random_integer(len(claves))]
	return planeta
}

// Función que envia un mensaje al servidor central
func send_message_to_server(capitan string, serverAddr string) {
	// Establecer conexión con el servidor
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	// Crear mensaje

	message := []byte(capitan + planet_asignation(100))

	// Enviar mensaje al servidor
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error al enviar mensaje al servidor:", err)
		return
	}
	fmt.Println("Capitán", capitan, "encontró botín en Planeta", string(message[1:]), ", enviando solicitud de asignación")
}

func main() {
	// Dirección del servidor
	serverAddr := "localhost:8080" // Reemplazar con la dirección del servidor y el puerto

	// Establecer conexión con el servidor
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	// Enviar mensaje al servidor
	send_message_to_server("1", serverAddr)
	send_message_to_server("2", serverAddr)
	send_message_to_server("3", serverAddr)
}
