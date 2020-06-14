package logger

import (
	"fmt"
	c "ph1-emulator/constants"
	"ph1-emulator/numbers"
)

// LogRTL converte uma operação para notação RTL e exibe
func LogRTL(opName string, value int) {
	// Procura na tabela de simbolos pelo nome da operação
	opSymbol := OperationSymbol[opName]
	// Converte para hexadecimal para imprimir corretamente
	hexValue := numbers.IntToHex(value)

	// Chama o handler que irá tratar cada RTL
	rltExpression := handleRtlExpression(opName, opSymbol, hexValue)

	// Exibe a expressão rtl já preenchida
	fmt.Println(rltExpression)
}

func handleRtlExpression(name string, opSymbol string, hexValue string) string {
	// Chama o método de agrupamento para facilitar na busca pelo RTL
	simplifiedName := Simplify(name)
	// Busca a expressão RTL usando o nome simplificado
	rtlExpression := GetRtlExpression[simplifiedName]
	var filledRtl string

	// Preenche o RTL com os valores necessários para exibição correta
	// de acordo com a operação, verificando quais usam o mesmo formato
	// de mensagem
	if simplifiedName == c.Ula || simplifiedName == c.CondOp {
		filledRtl = fmt.Sprintf(rtlExpression, name, hexValue, opSymbol, hexValue)
	} else if simplifiedName == c.NopOp || simplifiedName == c.HaltOp || simplifiedName == c.NotOp {
		filledRtl = fmt.Sprintf(rtlExpression)
	} else {
		filledRtl = fmt.Sprintf(rtlExpression, hexValue, hexValue)
	}
	return filledRtl
}
