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
	// Abre o arquivo como um bloco de string com todo o conteudo
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file: %s", fileName)
	}
	defer file.Close()

	var lines []string
	// Cria um novo Buffer para ler linha por linha do bloco
	scanner := bufio.NewScanner(file)
	// For de iteração, para cada linha atribui a uma variavel text removendo os espacos(se tiver)
	// do inicio e do fim da linha. Se houver conteudo atribui para a lista
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		if len(text) > 0 {
			lines = append(lines, text)
		}
	}
	// Retorna uma lista de string contendo todas as linhas do arquivo por ordem de entrada
	return lines
}

// MapInstructionsToMemory recebe uma lista de linhas contendo endereços de memória e instruções
// para então mapeá-las para a memória virtual
func MapInstructionsToMemory(instructions []string) {
	// le cada linha e mapeia o conteudo em endereco e  valor
	for _, instruction := range instructions {
		var address, value string

		// Pega o endereço de memória e a instrução de acordo com a posição no
		// arquivo. Exemplo: 02 00
		fmt.Sscanf(instruction, "%2s %2s", &address, &value)

		// Normaliza os valores para evitar erros na verificação abaixo
		address = strings.Trim(address, " ")
		value = strings.Trim(value, " ")

		// Se houver valores converte para uint16 e atribui
		// na memória virtual
		if len(address) == 2 && len(value) == 2 {
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
	// Chama o mapper de instruções onde verifica os valores e atribui na memória
	// passando como parâmetro a função de leitura do arquivo que retorna uma lista
	// de string contendo as linhas do arquivo
	MapInstructionsToMemory(ReadFileContent(fileName))
}
