package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ph1-emulator/control"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
)

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
func mapFileInfoToVirtualMemory(instructions []string) {
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
		memory.VirtualMemory.SetValue(addr, val)
	}
}

func main() {
	fileName := getFileName()
	instructionFile := readFile(fileName)

	mapFileInfoToVirtualMemory(instructionFile)

	control.UnityControl.Start()
	log.Print(control.UnityControl.Counter)
}
