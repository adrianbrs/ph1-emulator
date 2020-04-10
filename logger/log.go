package logger

import (
	"fmt"
	"ph1-emulator/memory"
	"ph1-emulator/regs"
)

// LogFinalState exibe os valores atuais dos registradores,
// e os valores atuais dos endereços modificados na memória
func LogFinalState() {
	// Log dos registradores
	fmt.Println()
	fmt.Println("Registers:")
	fmt.Printf("AC %02X\n", regs.RegAC.Value)
	fmt.Printf("PC %02X\n", regs.RegPC.GetValue())
	fmt.Println()

	// Log da memória
	fmt.Println("Memory:")

	// Pega todos os endereços que foram modificados após a memória ser carregada
	// e exibe os valores atuais correspondentes
	for _, addr := range memory.VirtualMemory.GetChangedAddresses() {
		fmt.Printf("%02X %02X\n", addr, memory.VirtualMemory.GetValue(addr))
	}
}
