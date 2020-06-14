package logger

import (
	c "ph1-emulator/constants"
)

// GetRtlExpression mapeia uma operação para seu respectivo RTL
var GetRtlExpression = map[string]string{
	c.Ula:    "%s %s ; AC <- AC %s MEM[%s]",
	c.CondOp: "%s %s ; Se AC%s0 então PC <- %s",
	c.NopOp:  "NOP",
	c.StrOp:  "STR %s ; MEM[%s] <- AC",
	c.NotOp:  "NOT    ; AC <- !AC",
	c.HaltOp: "HLT",
	c.JmpOp:  "JMP %s ; PC <- %s",
	c.LdrOp:  "LDR %s ; AC <- MEM[%s]",
}

// OperationSymbol Armazena os símbolos RTL de cada operação lógica e matemática
var OperationSymbol = map[string]string{
	c.AddOp: "+",
	c.SubOp: "-",
	c.MulOp: "*",
	c.DivOp: "/",
	c.AndOp: "&",
	c.OrOp:  "|",
	c.XorOp: "^",
	c.JeqOp: "==",
	c.JgOp:  ">",
	c.JlOp:  "<",
}

// Simplify agrupa as operações da ula e condicionais em nomes genéricos
// para facilitar a busca no map de RTL
func Simplify(name string) string {
	if name == c.AddOp || name == c.SubOp || name == c.DivOp || name == c.MulOp ||
		name == c.AndOp || name == c.OrOp || name == c.XorOp {
		name = c.Ula
	} else if name == c.JeqOp || name == c.JgOp || name == c.JlOp {
		name = c.CondOp
	}
	return name
}
