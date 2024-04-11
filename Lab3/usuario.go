package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random_integer(n int, m int) int {
	return n + rand.Intn(m-n)
}

func main() {
	start := time.Now().Second()
	for time.Now().Second()-start < 10 {
		// Esperar 7 segundos
		time.Sleep(70 * time.Millisecond)
	}

	AT := random_integer(20, 30)
	MP := random_integer(10, 15)
	acc := false
	for !acc {
		// Esperar 3 segundos
		time.Sleep(70 * time.Millisecond)
		acc = true // Consulta al servidor (borrar true al agregar consulta)

		fmt.Print("Solicitando", AT, "y", MP, "MP ; ")
		if acc {
			fmt.Println("Resolución: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicación")
		} else {
			fmt.Println("Resolución: -- DENEGADA -- ; Reintentando en 3 segs...")
		}

	}

}
