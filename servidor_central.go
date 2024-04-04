package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
)

func random_integer(n int) int {
	return rand.Intn(n)
}

func planet_asignation(n int) map[string]int {
	// Crear un mapa para almacenar los valores
	diccionario := make(map[string]int)

	// Definir las claves "PA", "PB", ..., "PF"
	claves := []string{"PA", "PB", "PC", "PD", "PE", "PF"}

	// Generar valores para cada clave
	for _, clave := range claves {
		diccionario[clave] = random_integer(n)
	}
	return diccionario
}

func less_loot(planets map[string]int, n int) string {
	min := n + 1
	minName := ""
	for key, v := range planets {
		if v < min {
			min = v
			minName = key
		}
	}
	return minName
}

func main() {
	// Port and buffer
	port := ":8080"
	buffer := 1024

	// Comprueba si el puerto está activo
	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Crea la conexión
	conexion, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Cierra la conexión
	defer conexion.Close()
	buffer_entrada := make([]byte, buffer)

	// Loot max value
	n := 100

	// Planets loot
	planets := planet_asignation(n)
	fmt.Println(planets)

	// Bucle infinito que escucha los mensajes, imprime el mensaje recibido y responden
	for {

		// Lee el mensaje
		n, addr, err := conexion.ReadFromUDP(buffer_entrada)
		// Imprime el mensaje
		fmt.Println("->", string(buffer_entrada[0:n]), "desde", addr)
		// Comprueba si el mensaje es STOP
		if strings.TrimSpace(string(buffer_entrada[0:n])) == "STOP" {
			fmt.Println("El servidor se ha detenido")
			return
		}

		// Responde al mensaje
		mensaje := []byte("Mensaje recibido")
		// Imprime el mensaje enviado
		fmt.Println("Mensaje enviado ->", mensaje)
		// Envia el mensaje
		_, err = conexion.WriteToUDP(mensaje, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
