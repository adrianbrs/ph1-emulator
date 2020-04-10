package memory

import (
	"log"
	"ph1-emulator/config"
	"strconv"
)

// virtualMemory representação virtual de uma memória
type virtualMemory struct {
	Data           map[int]int
	loaded         bool
	runtimeChanges []int
}

// SetValue define o valor de um endereço
func (mem *virtualMemory) SetValue(addr int, value int) {
	addr = parseAddr(addr)
	mem.Data[addr] = parseWord(value)

	// Adiciona o endereço na lista de endereços modificados
	// caso a memória já tenha sido carregada
	if mem.loaded {
		mem.runtimeChanges = append(mem.runtimeChanges, addr)
	}
}

// GetValue resgata um valor presente em um endereço
func (mem *virtualMemory) GetValue(addr int) int {
	return mem.Data[parseAddr(addr)]
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
func (mem *virtualMemory) GetChangedAddresses() []int {
	return mem.runtimeChanges
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

func parseAddr(addr int) int {
	res, err := strconv.ParseUint(strconv.Itoa(addr), 10, config.AddrLength)
	if err != nil {
		log.Fatalf("virtualMemory address overflow: %d", addr)
	}
	return int(res)
}

func parseWord(val int) int {
	res, err := strconv.ParseUint(strconv.Itoa(val), 10, config.WordLength)
	if err != nil {
		log.Fatalf("virtualMemory word length overflow: %d", val)
	}
	return int(res)
}
