package control

import (
	"log"
	"ph1-emulator/decoder"
	"ph1-emulator/memory"
	"ph1-emulator/regs"
)

// unityControl => Unidade de Controle
type unityControl struct {
	Houte   bool
	Counter int
}

// GetNextInstruction pega a proxima instrução na memória
func (uc *unityControl) GetNextInstruction() int {
	return memory.VirtualMemory.GetValue(regs.RegPC.GetValue())
}

// Start inicia a unidade de controle
func (uc *unityControl) Start() {
	for !uc.Houte {
		nextInstruction := uc.GetNextInstruction()
		name, hasAddress := decoder.DecodeInstruction(nextInstruction)

		// PC++
		regs.RegPC.Increment()

		var end int
		if hasAddress {
			// Endereço está na proxima instrução
			end = uc.GetNextInstruction()
			regs.RegPC.Increment()
		}

		log.Print("Nome: "+name+", ", end)

		if name == "HLT" {
			uc.Houte = true
		}
		uc.Counter++
	}
}

// NewUnityControl retorna uma nova instância da unidade de controle
func newUnityControl() *unityControl {
	return &unityControl{
		Houte:   false,
		Counter: 0,
	}
}

// UnityControl => Unidade de Controle
var UnityControl = newUnityControl()
