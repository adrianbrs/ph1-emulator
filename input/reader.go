package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ph1-emulator/config"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
	"strings"
)

// readFileName lê o input do usuario cujo conteudo eh o nome do arquivo de instrucoes
// a ser aberto para leitura
func readFileName(msg string) string {
	var fileName string
	fmt.Print(msg)
	fmt.Scanln(&fileName)
	fmt.Println()
	return fileName
}

// ReadFileContent lê os dados do arquivo cujo nome foi inserido pelo usuario
func ReadFileContent(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file: %s", fileName)
	}
	defer file.Close()

	//scaneia linha à linha o bloco de string lido e retorna uma lista de string com todas as linhas
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		if len(text) > 0 {
			lines = append(lines, text)
		}
	}
	return lines
}

// MapInstructionsToMemory recebe uma lista de linhas contendo endereços de memória e instruções
// para então mapeá-las para a memória virtual
func MapInstructionsToMemory(instructions []string) {
	// le cada linha e mapeia o conteudo em endereco e  valor
	for _, instruction := range instructions {
		var address, value string

		// Pega o endereço de memória e a instrução
		fmt.Sscanf(instruction, "%s %s", &address, &value)

		// Armazena o valor no endereço de memória
		if len(address) > 0 && len(value) > 0 {
			memory.VirtualMemory.SetValue(
				numbers.HexToInt(address, config.AddrLength),
				numbers.HexToInt(value, config.WordLength))
		}
	}

	// Define o estado da memória como carregada
	memory.VirtualMemory.SetLoaded()
}

// RequestInput lê o arquivo de entrada contendo as instruções e mapeia
// as instruções na memória
func RequestInput() {
	// Verifica se o nome do arquivo foi passado nos argumentos
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fileName = readFileName("Input file: ")
	}

	if fileName == "" {
		log.Fatal("File name must not be empty")
	}

	MapInstructionsToMemory(ReadFileContent(fileName))
}
