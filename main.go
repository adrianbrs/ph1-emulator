package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ph1-emulator/decoder"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
	"ph1-emulator/registrators"
)

// UnityControl => Unidade de Controle
type UnityControl struct {
	Memory  *memory.VirtualMemory
	AC      *registrators.RegAC
	PC      *registrators.RegPC
	Houte   bool
	Counter int
}

// GetNextInstruction pega a proxima instrução
func (uc *UnityControl) GetNextInstruction() int {
	return uc.Memory.GetValue(uc.PC.GetValue())
}

// Start inicia a unidade de controle
func (uc *UnityControl) Start() {
	for !uc.Houte {
		nextInstruction := uc.GetNextInstruction()
		name, hasAddress := decoder.DecodeInstruction(nextInstruction)

		// PC++
		uc.PC.Increment()

		var end int
		if hasAddress {
			// Endereço está na proxima instrução
			end = uc.GetNextInstruction()
			uc.PC.Increment()
		}

		log.Print("Nome: "+name+", ", end)

		if name == "HLT" {
			uc.Houte = true
		}
		uc.Counter++
	}
}

func mapFileInfoToVirtualMemory(instructions []string, virtualMemory *memory.VirtualMemory) {
	var values map[string]string

	for _, instruction := range instructions {
		runeAuxiliar := []rune(instruction)
		address := string(runeAuxiliar[0:2])
		value := string(runeAuxiliar[3:5])

		values = map[string]string{
			address: value}
	}

	for addr, val := range values {
		addr := numbers.HexToInt(addr, 8)
		val := numbers.HexToInt(val, 8)
		virtualMemory.SetValue(addr, val)
	}
}

func getFileName() string {
	var instruction string
	fmt.Println("Digite o nome do arquivo desejado:")
	fmt.Scan(&instruction)

	return instruction
}

func readFile(fileName string) []string {
	var instructions []string
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Unable to open file")
	}

	reader := bufio.NewReader(file)
	reader.ReadLine()

	return instructions
}

func main() {
	fileName := getFileName()
	instructionFile := readFile(fileName)
	virtualMemory := memory.New()

	mapFileInfoToVirtualMemory(instructionFile, virtualMemory)

	uc := &UnityControl{
		Memory:  virtualMemory,
		AC:      registrators.NewRegAC(),
		PC:      registrators.NewRegPC(),
		Houte:   false,
		Counter: 0,
	}

	uc.Start()
	log.Print(uc.Counter)
}
