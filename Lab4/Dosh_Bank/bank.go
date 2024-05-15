package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync" // Importar sync para manejar locks
)

var (
	mutex sync.Mutex // Definir un mutex para bloquear la escritura en el archivo
)

func createFile() {
	file, err := os.Create("dosh_bank.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	file.Close() // No se necesita defer aquí
}

func writeToFile(mercenary string, floor string, amount string) {
	file, err := os.OpenFile("dosh_bank.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()

	// Bloquear la escritura para garantizar el orden correcto
	mutex.Lock()
	defer mutex.Unlock()

	_, err = file.WriteString("- " + mercenary + " " + floor + " " + amount + "\n")
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func readFromFile() (int, error) {
	file, err := os.Open("dosh_bank.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var totalAmount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			amountStr := parts[2]
			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				return 0, err
			}
			totalAmount += amount
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return totalAmount, nil

}
