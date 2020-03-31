package memory

import (
	"log"
	"ph1-emulator/config"
	"strconv"
)

// VirtualMemory representação virtual de uma memória
type VirtualMemory struct {
	Data map[int]int
}

// SetValue define o valor de um endereço
func (mem *VirtualMemory) SetValue(addr int, value int) {
	mem.Data[parseAddr(addr)] = parseWord(value)
}

// GetValue resgata um valor presente em um endereço
func (mem *VirtualMemory) GetValue(addr int) int {
	return mem.Data[parseAddr(addr)]
}

// New retorna uma nova instância da memória virtual
func New() (mem *VirtualMemory) {
	mem = &VirtualMemory{
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
		log.Fatalf("VirtualMemory address overflow: %d", addr)
	}
	return int(res)
}

func parseWord(val int) int {
	res, err := strconv.ParseUint(strconv.Itoa(val), 10, config.WordLength)
	if err != nil {
		log.Fatalf("VirtualMemory word length overflow: %d", val)
	}
	return int(res)
}
