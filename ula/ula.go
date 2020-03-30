package ula

//ExecuteOperation executa a operação de acordo com o opcode
func ExecuteOperation(ulaOp string, memoryValue int, *ac int) {
	acInt8 := int8(ac)
	switch ulaOp {
	case "NOT":
		ac = ^pc
	case "ADD":
		ac += memoryValue
	case "SUM":
		ac -= memoryValue
	case "MUL":
		ac *= memoryValue
	case "DIV":
		pc /= memoryValue
	case "AND":
		acInt8 = acInt8 & memoryValue
		ac = int(pcInt8)
	case "OR":
		acInt8 = acInt8 | memoryValue
		ac = int(acInt8)
	case "XOR":
		acInt8 = acInt8 ^ memoryValue
		ac = int(acInt8)
	}
}
