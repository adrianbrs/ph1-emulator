package logger

import (
	"fmt"
	"ph1-emulator/numbers"
)

// Armazena os símbolos RTL de cada operação lógica e matemática
var operationSymbol = map[string]string{
	"ADD": "+",
	"SUB": "-",
	"MUL": "*",
	"DIV": "/",
	"AND": "&",
	"OR":  "|",
	"XOR": "^",
	"JEQ": "==",
	"JG":  ">",
	"JL":  "<",
}

// Lista com as operações válidas para a notação RTL
var operationList = []string{"NOP", "LDR", "STR", "ADD", "SUB", "MUL",
	"DIV", "NOT", "AND", "OR", "XOR", "JMP", "JEQ", "JG", "JL",
	"HLT"}

// LogRTL converte uma operação para notação RTL e exibe
func LogRTL(opName string, value int) {
	for index, name := range operationList {
		if opName == name {
			executeLogInformation(index, value, name)
		}
	}
}

// Exibe a notação RTL
func executeLogInformation(index int, value int, name string) {
	opSymbol := operationSymbol[name]
	hexValue := numbers.IntToHex(value)

	switch index {
	case 0:
		fmt.Printf("NOP %s\n", hexValue)
	case 1:
		fmt.Printf("LDR %s ; AC <- MEM[%s]\n", hexValue, hexValue)
	case 2:
		fmt.Printf("STR %s ; MEM[%s] <- AC\n", hexValue, hexValue)
	case 3, 4, 5, 6:
		fmt.Printf("%s %s ; AC <- AC %s MEM[%s]\n", name, hexValue, opSymbol, hexValue)
	case 7:
		fmt.Printf("NOT %s ; AC <- !AC\n", hexValue)
	case 8, 9, 10:
		fmt.Printf("%s %s ; AC <- AC %s MEM[%s]\n", name, hexValue, opSymbol, hexValue)
	case 11:
		fmt.Printf("JMP %s ; PC <- %s\n", hexValue, hexValue)
	case 12, 13, 14:
		fmt.Printf("%s %s ; Se AC%s0 então PC <- %s\n", name, hexValue, opSymbol, hexValue)
	case 15:
		fmt.Println("HLT")
	}

}