package ula

//ExecuteOperation executa a operação de acordo com o opcode
func ExecuteOperation(ulaOp string, memoryValue int, ac int) int {
	acInt8 := int8(ac)
	memoryValueInt8 := int8(memoryValue)
	switch ulaOp {
	case "NOT":
		ac = ^ac
	case "ADD":
		ac += memoryValueInt8
	case "SUM":
		ac -= memoryValueInt8
	case "MUL":
		ac *= memoryValueInt8
	case "DIV":
		ac /= memoryValueInt8
	case "AND":
		acInt8 = acInt8 & memoryValueInt8
		ac = int(acInt8)
	case "OR":
		acInt8 = acInt8 | memoryValueInt8
		ac = int(acInt8)
	case "XOR":
		acInt8 = acInt8 ^ memoryValueInt8
		ac = int(acInt8)
	}
	return ac
}
