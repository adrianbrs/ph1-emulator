package logger

import (
	"fmt"
	"ph1-emulator/numbers"
)

// LogRTL converte uma operação para notação RTL e exibe
func LogRTL(opName string, value int) {
	executeLogInformation(value, opName)
}

// Exibe a notação RTL
func executeLogInformation(value int, name string) {
	// Procura na tabela de simbolos pelo nome da operação
	opSymbol := OperationSymbol[name]
	// Converte para hexadecimal para imprimir corretamente
	hexValue := numbers.IntToHex(value)

	// Chama o handler que irá tratar cada RTL
	rltExpression := handleRtlExpression(name, opSymbol, hexValue)

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
	if simplifiedName == "ULA" || simplifiedName == "COND" {
		filledRtl = fmt.Sprintf(rtlExpression, name, hexValue, opSymbol, hexValue)
	} else if simplifiedName == "NOP" || simplifiedName == "HLT" || simplifiedName == "NOT" {
		filledRtl = fmt.Sprintf(rtlExpression)
	} else {
		filledRtl = fmt.Sprintf(rtlExpression, hexValue, hexValue)
	}
	return filledRtl
}
