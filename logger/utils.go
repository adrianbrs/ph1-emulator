package logger

// GetRtlExpression mapeia uma operação para seu respectivo RTL
var GetRtlExpression = map[string]string{
	"ULA":  "%s %s ; AC <- AC %s MEM[%s]",
	"COND": "%s %s ; Se AC%0 então PC <- %s",
	"NOP":  "NOP",
	"STR":  "STR %s ; MEM[%s] <- AC",
	"NOT":  "NOT    ; AC <- !AC",
	"HLT":  "HLT",
	"JMP":  "JMP %s ; PC <- %s",
	"LDR":  "LDR %s ; AC <- MEM[%s]",
}

// OperationSymbol Armazena os símbolos RTL de cada operação lógica e matemática
var OperationSymbol = map[string]string{
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

// Simplify agrupa as operações da ula e condicionais em nomes genéricos
// para facilitar a busca no map de RTL
func Simplify(name string) string {
	if name == "ADD" || name == "SUB" || name == "DIV" || name == "MUL" ||
		name == "AND" || name == "OR" || name == "XOR" {
		name = "ULA"
	} else if name == "JEQ" || name == "JG" || name == "JL" {
		name = "COND"
	}
	return name
}
