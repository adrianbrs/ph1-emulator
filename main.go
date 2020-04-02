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

//getFileName le o input do usuario cujo conteudo eh o nome do arquivo de instrucoes
//a ser aberto para leitura
func getFileName() string {
	var instruction string
	fmt.Println("Digite o nome do arquivo desejado:")
	fmt.Scan(&instruction)

	return instruction
}

//readFile lê o arquivo cujo nome foi inserido pelo usuario na funcao getFileName()
//scaneia linha à linha o bloco de string lido e retorna uma listra de string com todas as linhas
func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

//Recebe a lista de linhas lidas do arquivo e uma instancia de memoria virtual
func mapFileInfoToVirtualMemory(instructions []string, virtualMemory *memory.VirtualMemory) {
	var values = map[string]string{}

	//le cada linha e mapeia o conteudo em endereco e  valor
	for _, instruction := range instructions {
		address := instruction[0:2]
		value := instruction[3:5]
		values[address] = value
	}

	//recebe esse map e converte em inteiro para popular os campos da memoria virtual
	for addr, val := range values {
		addr := numbers.HexToInt(addr, 8)
		val := numbers.HexToInt(val, 8)
		virtualMemory.SetValue(addr, val)
	}
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
