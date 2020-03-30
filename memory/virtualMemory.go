package memory

import (
	"fmt"
	"log"
	"ph1-emulator/config"
	"reflect"
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
		Data: make(map[int]int)
	}

	// Preenche a memória com 0
	for idx = 0; idx < 256; idx++ {
		mem.Data[idx] = 0
	}

	return mem
}

func parseAddr(addr int) int {
	addr, err := strconv.ParseUint(strconv.Itoa(addr), config.HexaBaseValue, config.AddrLength)
	if err != nil {
		log.Fatal("VirtualMemory address overflow: %d", addr)
	}
	return int(addr)
}

func parseWord(val int) int {
	res, err := strconv.ParseUint(strconv.Itoa(val), config.HexaBaseValue, config.WordLength)
	if err != nil {
		log.Fatal("VirtualMemory word length overflow: %d", val)
	}
	return int(res)
}
