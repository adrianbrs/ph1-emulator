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
	Data map[uint]uint
}

// SetValue define o valor de um endereço
func (mem *VirtualMemory) SetValue(addr int, value int) {
	mem.Data[parseAddr(addr)] = parseWord(value)
}

// GetValue resgata um valor presente em um endereço
func (mem *VirtualMemory) GetValue(addr int) uint8 {
	return mem.Data[parseAddr(addr)]
}

// New retorna uma nova instância da memória virtual
func New() (mem *VirtualMemory) {
	mem = &VirtualMemory{
		Data: make(map[uint]uint)
	}

	// Preenche a memória com 0
	for idx = 0; idx < 256; idx++ {
		mem.Data[idx] = 0
	}

	return mem
}

func parseAddr(addr int) uint8 {
	addr, err := strconv.ParseUint(strconv.Itoa(addr), config.HexaBaseValue, config.AddrLength)
	if err != nil {
		log.Panicf("Could not use '%s' as a valid memory address.")
	}
	return addr
}
