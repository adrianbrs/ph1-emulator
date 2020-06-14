package decoder

import (
	"log"
	c "ph1-emulator/constants"
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
	"00": {name: c.NopOp, hasAddress: false},
	"10": {name: c.LdrOp, hasAddress: true},
	"20": {name: c.StrOp, hasAddress: true},
	"30": {name: c.AddOp, hasAddress: true},
	"40": {name: c.SubOp, hasAddress: true},
	"50": {name: c.MulOp, hasAddress: true},
	"60": {name: c.DivOp, hasAddress: true},
	"70": {name: c.NotOp, hasAddress: false},
	"80": {name: c.AndOp, hasAddress: true},
	"90": {name: c.OrOp, hasAddress: true},
	"A0": {name: c.XorOp, hasAddress: true},
	"B0": {name: c.JmpOp, hasAddress: true},
	"C0": {name: c.JeqOp, hasAddress: true},
	"D0": {name: c.JgOp, hasAddress: true},
	"E0": {name: c.JlOp, hasAddress: true},
	"F0": {name: c.HaltOp, hasAddress: false},
}

// DecodeInstruction decodifica uma instrução pelo seu opcode
func DecodeInstruction(opcodeInt int) (string, bool) {
	opcodeHex := numbers.IntToHex(opcodeInt)
	instruction := availableInstructions[opcodeHex]

	if instruction.name == c.Empty {
		log.Fatalf(c.InvalidInstruction, opcodeHex)
	}
	return instruction.name, instruction.hasAddress
}
