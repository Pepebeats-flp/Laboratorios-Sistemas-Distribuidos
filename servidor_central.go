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

// Función que retorna el planeta el nombre del planeta con menos botín
func less_loot(planets map[string]int, n int) string {
	// Definir el valor máximo
	min := n

	// Definir el planeta con menos botín
	planeta_min := ""

	// Recorrer el mapa
	for planeta, botin := range planets {
		// Comprobar si el botín es menor que el mínimo
		if botin < min {
			// Actualizar el mínimo
			min = botin
			// Actualizar el planeta con menos botín
			planeta_min = planeta
		}
	}
	fmt.Println("Planeta con menos botín:", planeta_min)
	return planeta_min
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

	// Bucle infinito que escucha los mensajes, imprime el mensaje recibido y responden
	for {

		// Lee el mensaje
		n, addr, err := conexion.ReadFromUDP(buffer_entrada)

		// Imprime el mensaje
		fmt.Println("Recepción de solicitud desde el Planeta", string(buffer_entrada[1:3]), ", del capitán", string(buffer_entrada[0:1]))

		// Comprueba si el mensaje es STOP
		if strings.TrimSpace(string(buffer_entrada[0:n])) == "STOP" {
			fmt.Println("El servidor se ha detenido")
			return
		}

		// Imprimimos estado actual de los planetas
		fmt.Println("Estado actual de los planetas:" + fmt.Sprint(planets))

		// Obtiene el planeta con menos botín
		planeta_min := less_loot(planets, n)
		// Le sumamos 1 al botín del planeta_min
		planets[planeta_min] = planets[planeta_min] + 1
		// Imprime el planeta asignado (Bot ́ın asignado al planeta PC, cantidad actual: 5)
		fmt.Println("Botín asignado al planeta", planeta_min, ", cantidad actual:", planets[planeta_min])

		// Responde al mensaje (no se si lo ocupemos xd)
		mensaje := []byte("Mensaje recibido")
		// Imprime el mensaje enviado
		//fmt.Println("Mensaje enviado ->", mensaje)
		// Envia el mensaje
		_, err = conexion.WriteToUDP(mensaje, addr)
		if err != nil {
			//fmt.Println(err)
			return
		}
	}
}
