package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
)

// readFileName lê o input do usuario cujo conteudo eh o nome do arquivo de instrucoes
// a ser aberto para leitura
func readFileName(msg string) string {
	var fileName string
	fmt.Print(msg)
	fmt.Scan(&fileName)
	fmt.Println()
	return fileName
}

// readFileContent lê os dados do arquivo cujo nome foi inserido pelo usuario
func readFileContent(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file: %s", fileName)
	}
	defer file.Close()

	//scaneia linha à linha o bloco de string lido e retorna uma lista de string com todas as linhas
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Recebe uma lista de linhas contendo endereços de memória e instruções
// para então mapeá-las para a memória virtual
func mapInstructionsToMemory(instructions []string) {
	// le cada linha e mapeia o conteudo em endereco e  valor
	for _, instruction := range instructions {
		// Pega o endereço de memória em que a instrução vai ser guardada
		address := numbers.HexToInt(instruction[0:2], 8)

		// Pega o valor, em hexadecimal, correspondente a uma instrução
		value := numbers.HexToInt(instruction[3:5], 8)

		// Armazena o valor no endereço de memória
		memory.VirtualMemory.SetValue(address, value)
	}

	// Define o estado da memória como carregada
	memory.VirtualMemory.SetLoaded()
}

// ReadInstructionsFile lê o arquivo de entrada contendo as instruções e mapeia
// as instruções na memória
func ReadInstructionsFile(inputMessage string) {
	fileName := readFileName(inputMessage)
	mapInstructionsToMemory(readFileContent(fileName))
}
