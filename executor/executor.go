package executor

import (
	"reflect"
	"strings"
)

// ExecuteOperation executa a operação na ula/uc de acordo com seu nome
func ExecuteOperation(obj interface{}, name string, args ...interface{}) bool {
	// Por meio de reflection recebe uma interface(um endereço de ULA ou de UC)
	// e concatena o nome da operação recebido por parâmetro com o padrão dos métodos,
	// conseguindo chamar qualquer um dos métodos de ambas as interfaces com a mesma
	// linha de código
	method := reflect.ValueOf(obj).MethodByName("ExecuteOp" + strings.ToUpper(name))

	// Verifica se o método resultante da chamada é válido
	if !method.IsValid() {
		// Retorna falso pois nenhuma operação aconteceu
		return false
	}

	// Cria por meio do make um slice de valores que servirá como
	// parâmetro para a chamada do método. Esse slice é do tipo
	// reflect.Value(tipo dinâmico) e é de tamanho dinâmico de acordo com
	// a quantidade de argumentos encontrados na variável method
	operationArgs := make([]reflect.Value, method.Type().NumIn())

	// Itera atribuindo os valores recebidos por parâmetro
	// (podem ser infinitos parâmetros, visto que foi usado sinal ...)
	// atribuindo para o indice i da lista de argumentos da operação
	for i := 0; i < method.Type().NumIn(); i++ {
		operationArgs[i] = reflect.ValueOf(args[i])
	}

	// Chama o método contigo na variável method(encontrada no inicio da função
	// por meio de reflection) passando a lista de argumentos
	method.Call(operationArgs)

	// Retorna true pois a operação aconteceu
	return true
}
