package control

import (
	"log"
	c "ph1-emulator/constants"
	"ph1-emulator/decoder"
	"ph1-emulator/executor"
	"ph1-emulator/logger"
	"ph1-emulator/memory"
	"ph1-emulator/regs"
	"ph1-emulator/ula"
)

// unityControl => Virtualização da Unidade de Controle
type unityControl struct {
	Houte   bool
	Counter int64
}

// GetNextInstruction pega a proxima instrução na memória
func (uc *unityControl) GetNextInstruction() int {
	return memory.VirtualMemory.GetValue(regs.RegPC.GetValue())
}

// Execute tenta encontrar e executar a função correspondente
// ao mnemônico de uma instrução
func (uc *unityControl) Execute(opName string, value int) {
	found := executor.ExecuteOperation(uc, opName, value) ||
		executor.ExecuteOperation(ula.GetInstance(), opName, value)

	// Verifica se a função da instrução não foi encontrada
	// e para a execução retornando um erro
	if !found {
		log.Fatalf(c.OperationNotFound, opName)
	}

}

// Start inicia a unidade de controle
func (uc *unityControl) Start(halt chan bool) {
	// Loop enquanto não encontrar a operação de parada (HTL)
	for !uc.Houte {

		// Pega a próxima instrução a ser executada
		nextInstruction := uc.GetNextInstruction()

		// Decodifica a instrução encontrada
		name, hasAddress := decoder.DecodeInstruction(nextInstruction)

		// Incrementa o Program Counter
		regs.RegPC.Increment()

		// Verifica se há endereço e pega através
		// do valor da próxima instrução
		var end int
		if hasAddress {
			// Endereço está na proxima instrução
			end = uc.GetNextInstruction()
			// Incrementa o Program Counter novamente pois essa instrucao
			// usou dois bits
			regs.RegPC.Increment()
		}

		// Executa a instrução encontrada após incrementar PC
		uc.Execute(name, end)

		// Aumenta o contador de insturções executadas
		uc.Counter++

		// Exibe a notação RTL da operação atual
		logger.LogRTL(name, end)
	}

	// Envia o sinal de término de execução (HLT)
	halt <- true
}

// UnityControl => Unidade de Controle
var UnityControl = &unityControl{
	Houte:   false,
	Counter: 0,
}
