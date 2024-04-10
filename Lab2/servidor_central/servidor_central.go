package main

import (
	"fmt"
	"math/rand"
	"net"
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
	min := n + 1

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
		} else if botin == min && planeta < planeta_min {
			planeta_min = planeta
		}
	}
	fmt.Println("Planeta con menos botín:", planeta_min)
	return planeta_min
}

func verify_done(planets map[string]int) bool {
	claves := []string{"PA", "PB", "PC", "PD", "PE", "PF"}
	for i := 1; i < len(claves); i++ {
		if planets["PA"] != planets[claves[i]] {
			return false
		}
	}
	return true
}

func main() {
	// Port and buffer
	port := ":8080"
	buffer := 1024

	// Comprueba si el puerto está activo
	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0"+port)
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
	n := 50

	// Planets loot
	planets := planet_asignation(n)

	// Bucle infinito que escucha los mensajes, imprime el mensaje recibido y responden
	for {
		// Lee el mensaje
		_, addr, _ := conexion.ReadFromUDP(buffer_entrada)

		// Imprime el mensaje
		fmt.Println("Recepción de solicitud desde el Planeta", string(buffer_entrada[1:]), ", del capitán", string(buffer_entrada[0:1]), "\n")

		// Imprimimos estado actual de los planetas
		fmt.Println("Estado actual de los planetas:")
		fmt.Println("PA:", planets["PA"], "PB:", planets["PB"], "PC:", planets["PC"], "PD:", planets["PD"], "PE:", planets["PE"], "PF:", planets["PF"], "\n")

		// Obtiene el planeta con menos botín
		planeta_min := less_loot(planets, n)
		// Le sumamos 1 al botín del planeta_min
		planets[planeta_min] = planets[planeta_min] + 1
		// Imprime el planeta asignado (Botín asignado al planeta PC, cantidad actual: 5)
		fmt.Println("Botín asignado al planeta", planeta_min, ", cantidad actual:", planets[planeta_min], "\n")

		// Verifica si todos los planetas tienen la misma cantidad de botín
		if verify_done(planets) {
			fmt.Println("Todos los planetas tienen la misma cantidad de botín.")
			fmt.Println("Notificando a los capitanes...")

			mensaje := []byte("Stop")
			// Envia el mensaje
			_, err = conexion.WriteToUDP(mensaje, addr)
			if err != nil {
				//fmt.Println(err)
				return
			}

			fmt.Println("El servidor se ha detenido.")
			return
		} else {
			mensaje := []byte("Listening...")
			// Envia el mensaje
			_, err = conexion.WriteToUDP(mensaje, addr)
			if err != nil {
				//fmt.Println(err)
				return
			}
		}

	}
}
