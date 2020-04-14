package memory

import (
	"log"
	"ph1-emulator/numbers"
)

// virtualMemory representação virtual de uma memória
type virtualMemory struct {
	Data           map[int]int
	loaded         bool
	runtimeChanges []int
}

// SetValue define o valor de um endereço
func (mem *virtualMemory) SetValue(addr int, value int) {
	// Faz o parse do valor do endereco para inteiro de 8 bits positivo(uint8)
	addr = parseAddr(addr)

	// Verifica se a memoria ja foi carregada
	if mem.loaded {
		// Verifica se o endereco nao eh uma instrucao
		if addr < 80 {
			log.Fatal("Can't change instruction memory area during runtime")
		}
		// Adiciona o valor à lista de valores modificados da memória
		mem.runtimeChanges = append(mem.runtimeChanges, addr)
	}

	// Seta o valor na memoria
	mem.Data[addr] = parseWord(value)
}

// GetValue resgata um valor presente em um endereço
func (mem *virtualMemory) GetValue(addr int) int {
	return mem.Data[parseAddr(addr)]
}

// GetWordValue retorna o valor presente no endereço especificado como um
// inteiro convertido para int8 e depois int
func (mem *virtualMemory) GetWordValue(addr int) int {
	return int(int8(mem.GetValue(addr)))
}

// SetLoaded define que a memória virtual foi carregada com sucesso
func (mem *virtualMemory) SetLoaded() {
	mem.loaded = true
}

// IsLoaded retorna se os valores do programa ja foram carregados
func (mem *virtualMemory) IsLoaded() bool {
	return mem.loaded
}

// GetChangedAddresses retorna uma lista com os endereços de memória
// que forma modificados
func (mem *virtualMemory) GetLastChanges() map[int]int {
	changes := make(map[int]int, len(mem.runtimeChanges))
	for _, addr := range mem.runtimeChanges {
		changes[addr] = mem.GetWordValue(addr)
	}
	return changes
}

// VirtualMemory => Memória Virtual
var VirtualMemory = newMemory()

// New retorna uma nova instância da memória virtual
func newMemory() (mem *virtualMemory) {
	mem = &virtualMemory{
		Data: make(map[int]int),
		// changedAddr: []int{},
		loaded: false,
	}

	// Preenche a memória com 0
	for idx := 0; idx < 256; idx++ {
		mem.Data[idx] = 0
	}

	return mem
}

// Faz o parse do valor do endereco de memoria pra inteiro de 8 bits
// positivo(uint8) e valida se houve overflow numerico
func parseAddr(addr int) int {
	res, err := numbers.ValidateAddress(addr)
	if err != nil {
		log.Fatalf("virtualMemory address overflow: %d", addr)
	}
	return res
}

// Faz o parse da palavra(instrucao) para inteiro de 8 bits
// positivo(uint8) e valida se houve overflow numerico
func parseWord(val int) int {
	res, err := numbers.ValidateWord(val)

	if err != nil {
		log.Fatalf("virtualMemory word length overflow: %d", val)
	}
	return res
}
