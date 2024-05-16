package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Estructura para el registro de decisiones
type DecisionRecord struct {
	Mercenary  string
	Floor      string
	IPDatanode string
	Content    string
}

// Estructura de NameNode
type NameNode struct {
	mu        sync.Mutex
	records   []DecisionRecord
	dataNodes []string
}

// NewNameNode crea una nueva instancia de NameNode con los DataNodes disponibles
func NewNameNode(dataNodes []string) *NameNode {
	return &NameNode{
		records:   make([]DecisionRecord, 0),
		dataNodes: dataNodes,
	}
}

// RecordDecision registra una decisión en un DataNode aleatorio y la guarda en el archivo decisions.txt
func (nn *NameNode) RecordDecision(record DecisionRecord) {
	nn.mu.Lock()
	defer nn.mu.Unlock()

	// Asignar un DataNode aleatorio
	rand.Seed(time.Now().UnixNano())
	record.IPDatanode = nn.dataNodes[rand.Intn(len(nn.dataNodes))]

	// Agregar el registro a la lista
	nn.records = append(nn.records, record)
	log.Printf("Decision recorded: %v", record)

	// Guardar el registro en un archivo
	nn.saveRecordToFile(record)

	// Enviar Mercenario,Piso y contenido por GRPC al DatNode de la IP correspondiente.
	// (no implementado en este ejemplo)
}

// GetDecisionRecords devuelve todas las decisiones registradas
func (nn *NameNode) GetDecisionRecords() []DecisionRecord {
	nn.mu.Lock()
	defer nn.mu.Unlock()

	return nn.records
}

// saveRecordToFile guarda el registro en un archivo de texto
func (nn *NameNode) saveRecordToFile(record DecisionRecord) {
	file, err := os.OpenFile("decisions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	line := fmt.Sprintf("%s Piso %s %s\n", record.Mercenary, record.Floor, record.IPDatanode)
	if _, err := file.WriteString(line); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func main() {
	// DataNodes disponibles
	dataNodes := []string{"10.0.0.15", "10.0.0.16", "10.0.0.17"}

	// Crear una instancia de NameNode
	nameNode := NewNameNode(dataNodes)

	// Ejemplo de registrar decisiones
	decision1 := DecisionRecord{
		Mercenary: "Mercenario1",
		Floor:     "1",
		Content:   "3",
	}
	decision2 := DecisionRecord{
		Mercenary: "Mercenario2",
		Floor:     "2",
		Content:   "A",
	}

	nameNode.RecordDecision(decision1)
	nameNode.RecordDecision(decision2)

	// Ejemplo de obtener todas las decisiones registradas
	records := nameNode.GetDecisionRecords()
	for _, record := range records {
		fmt.Printf("Mercenario: %s, Piso: %s, Contenido: %s, IP DataNode: %s\n", record.Mercenary, record.Floor, record.Content, record.IPDatanode)
	}
}
