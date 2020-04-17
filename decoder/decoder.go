package decoder

import (
	"log"
	"ph1-emulator/numbers"
)

// PH1Instruction constructor contendo mnemonico e se uma instrução
// necessita de um endereço
type PH1Instruction struct {
	name       string
	hasAddress bool
}

// Mapper de valor hexadecimal para o opCode(e se há endereço) respectivo
var availableInstructions = map[string]PH1Instruction{
	"00": {name: "NOP", hasAddress: false},
	"10": {name: "LDR", hasAddress: true},
	"20": {name: "STR", hasAddress: true},
	"30": {name: "ADD", hasAddress: true},
	"40": {name: "SUB", hasAddress: true},
	"50": {name: "MUL", hasAddress: true},
	"60": {name: "DIV", hasAddress: true},
	"70": {name: "NOT", hasAddress: false},
	"80": {name: "AND", hasAddress: true},
	"90": {name: "OR", hasAddress: true},
	"A0": {name: "XOR", hasAddress: true},
	"B0": {name: "JMP", hasAddress: true},
	"C0": {name: "JEQ", hasAddress: true},
	"D0": {name: "JG", hasAddress: true},
	"E0": {name: "JL", hasAddress: true},
	"F0": {name: "HLT", hasAddress: false},
}

// DecodeInstruction decodifica uma instrução pelo seu opcode
func DecodeInstruction(opcodeInt int) (string, bool) {
	opcodeHex := numbers.IntToHex(opcodeInt)
	instruction := availableInstructions[opcodeHex]

	if instruction.name == "" {
		log.Fatalf("Invalid instruction: none found for opCode %s", opcodeHex)
	}
	return instruction.name, instruction.hasAddress
}
