package logger

import (
	"fmt"
	"ph1-emulator/memory"
	"ph1-emulator/numbers"
	"ph1-emulator/regs"
	"sort"
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
	// Ordena os endereços
	changes := memory.VirtualMemory.GetLastChanges()
	addresses := make([]int, 0, len(changes))
	for addr := range changes {
		addresses = append(addresses, addr)
	}
	sort.Ints(addresses)

	for _, addr := range addresses {
		fmt.Printf("%s %s\n", numbers.IntToHex(addr), numbers.IntToHex(changes[addr]))
	}
}
