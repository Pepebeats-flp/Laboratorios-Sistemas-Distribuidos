package main

import (
	"fmt"
	"net"
)

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

	// Mensaje a enviar al servidor
	message := []byte("Mensaje de prueba")

	// Enviar mensaje al servidor
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error al enviar mensaje al servidor:", err)
		return
	}

	fmt.Println("Mensaje enviado al servidor:", string(message))
}
