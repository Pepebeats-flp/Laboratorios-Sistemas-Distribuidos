package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func random_integer(n int) int {
	return rand.Intn(n)
}

// Funcion que retorne un planeta al azar
func planet_looted(n int) string {
	// Definir las claves "PA", "PB", ..., "PF"
	claves := []string{"PA", "PB", "PC", "PD", "PE", "PF"}
	// Elegir un planeta al azar y retornarlo
	planeta := claves[random_integer(len(claves))]
	return planeta
}

// Función que envia un mensaje al servidor central
func send_message_to_server(capitan string, serverAddr string) bool {
	// Establecer conexión con el servidor
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return false
	}
	defer conn.Close()

	// Crear mensaje
	message := []byte(capitan + ":" + planet_looted(100))

	// Enviar mensaje al servidor
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error al enviar mensaje al servidor:", err)
		return false
	}
	fmt.Println("Capitán", capitan, "encontró botín en Planeta:", string(message[2:]), ", enviando solicitud de asignación")

	// Recibir mensaje del servidor
	buffer_entrada := make([]byte, 1024)
	_, err = conn.Read(buffer_entrada)
	if err != nil {
		fmt.Println("Error al recibir mensaje del servidor:", err)
		return false
	}
	response := strings.TrimRight(string(buffer_entrada), "\x00")
	if response == "Stop" {
		return false
	}
	return true
}

func main() {
	// Dirección del servidor
	serverAddr := "0.0.0.0:8080" // Reemplazar con la dirección del servidor y el puerto

	// Establecer conexión con el servidor
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	// Iniciar viaje capitanes
	for {
		time.Sleep(time.Duration(random_integer(3000)) * time.Millisecond)
		captain := random_integer(3) + 1
		if !send_message_to_server(strconv.Itoa(captain), serverAddr) {
			return
		}
	}
}
