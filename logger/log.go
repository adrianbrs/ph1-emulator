package logger

import (
	"fmt"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
	"ph1-emulator/regs"
)

// LogFinalState exibe os valores atuais dos registradores,
// e os valores atuais dos endereços modificados na memória
func LogFinalState(executionCount int) {
	// Log do número de execuções
	fmt.Println()
	fmt.Printf("%d instructions executed\n", executionCount)

	// Log dos registradores
	fmt.Println()
	fmt.Println("Registers:")
	fmt.Printf("AC %s\n", numbers.IntToHex(regs.RegAC.GetValue()))
	fmt.Printf("PC %s\n", numbers.IntToHex(regs.RegPC.GetValue()))
	fmt.Println()

	// Log da memória
	fmt.Println("Memory:")

	// Pega todos os endereços que foram modificados após a memória ser carregada
	// e exibe os valores atuais correspondentes
	for addr, val := range memory.VirtualMemory.GetLastChanges() {
		fmt.Printf("%s %s\n", numbers.IntToHex(addr), numbers.IntToHex(val))
	}
}
