package memory

import (
	"log"
	"ph1-emulator/config"
	"strconv"
)

// virtualMemory representação virtual de uma memória
type virtualMemory struct {
	Data map[int]int
}

// SetValue define o valor de um endereço
func (mem *virtualMemory) SetValue(addr int, value int) {
	mem.Data[parseAddr(addr)] = parseWord(value)
}

// GetValue resgata um valor presente em um endereço
func (mem *virtualMemory) GetValue(addr int) int {
	return mem.Data[parseAddr(addr)]
}

// VirtualMemory => Memória Virtual
var VirtualMemory = newMemory()

// New retorna uma nova instância da memória virtual
func newMemory() (mem *virtualMemory) {
	mem = &virtualMemory{
		Data: make(map[int]int),
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
