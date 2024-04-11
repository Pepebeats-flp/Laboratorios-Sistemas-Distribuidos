package main

import "time"

func main() {
	AT := 0
	MP := 0
	// Hacer lo de gRPC para recibir las consultas
	for true {
		// Esperar 5 segundos
		time.Sleep(50 * time.Millisecond)

		// Generar AT
		if AT < 40 {
			AT += 10
		} else {
			AT = 50
		}

		// Generar MP
		if MP < 15 {
			MP += 5
		} else {
			MP = 20
		}
	}
}

// Prints para agregar luego

// fmt.Print("Recepción de solicitud desde equipo", equipo, at_solicitado, "AT", mp_solicitado)
// fmt.Print(" -- APROBADA -- ")
// fmt.Print(" -- DENEGADA -- ")
// fmt.Println("AT EN SISTEMA:",AT,"; MP EN SISTEMA:", MP)
