package main

import (
	"log"
	"ph1-emulator/decoder"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
	"ph1-emulator/registrators"
)

// UnityControl => Unidade de Controle
type UnityControl struct {
	Memory  *memory.VirtualMemory
	AC      *registrators.RegAC
	PC      *registrators.RegPC
	Houte   bool
	Counter int
}

// GetNextInstruction pega a proxima instrução
func (uc *UnityControl) GetNextInstruction() int {
	return uc.Memory.GetValue(uc.PC.GetValue())
}

// Start inicia a unidade de controle
func (uc *UnityControl) Start() {
	for !uc.Houte {
		nextInstruction := uc.GetNextInstruction()
		name, hasAddress := decoder.DecodeInstruction(nextInstruction)

		// PC++
		uc.PC.Increment()

		var end int
		if hasAddress {
			// Endereço está na proxima instrução
			end = uc.GetNextInstruction()
			uc.PC.Increment()
		}

		log.Print("Nome: "+name+", ", end)

		if name == "HLT" {
			uc.Houte = true
		}
		uc.Counter++
	}
}

func main() {
	virtualMemory := memory.New()

	// Fake data
	fakeValues := map[string]string{
		"00": "10",
		"01": "81",
		"02": "30",
		"03": "82",
		"04": "20",
		"05": "80",
		"06": "F0",
		"80": "00",
		"81": "05",
		"82": "02",
	}
	for addr, val := range fakeValues {
		addr := numbers.HexToInt(addr, 8)
		val := numbers.HexToInt(val, 8)
		virtualMemory.SetValue(addr, val)
	}

	uc := &UnityControl{
		Memory:  virtualMemory,
		AC:      registrators.NewRegAC(),
		PC:      registrators.NewRegPC(),
		Houte:   false,
		Counter: 0,
	}

	uc.Start()

	log.Print(uc.Counter)
}
