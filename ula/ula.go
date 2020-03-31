package ula

//ExecuteOperation executa a operação de acordo com o opcode
func ExecuteOperation(ulaOp string, memoryValue int, *pc int) {
	pcInt8 := int8(pc)
	switch ulaOp {
	case "NOT":
		pc = ^pc
	case "ADD":
		pc += memoryValue
	case "SUM":
		pc -= memoryValue
	case "MUL":
		pc *= memoryValue
	case "DIV":
		pc /= memoryValue
	case "AND":
		pcInt8 = pcInt8 & memoryValue
		pc = int(pcInt8)
	case "OR":
		pcInt8 = pcInt8 | memoryValue
		pc = int(pcInt8)
	case "XOR":
		pcInt8 = pcInt8 ^ memoryValue
		pc = int(pcInt8)
	}
}
