package ula

import (
	"log"
	"reflect"
	"strings"
)

type ULA struct{}

func New() *ULA {
	ula := &ULA{}
	return ula
}

//ExecuteOperation executa a operação de acordo com o opcode
func (ula *ULA) Execute(ulaOp string, memoryValue int) {
	method := reflect.ValueOf(ula).MethodByName("ExecuteOp" + strings.ToUpper(ulaOp))

	if !method.IsValid() {
		log.Fatalf("ULA Operation %s not found", ulaOp)
	}

	operationArgs := make([]reflect.Value, method.Type().NumIn())

	if method.Type().NumIn() == 1 {
		operationArgs[0] = reflect.ValueOf(memoryValue)
	}
	method.Call(operationArgs)

}

func (ula *ULA) ExecuteOpNOT() {
	Ac.Value = ^Ac.Value
}

func (ula *ULA) ExecuteOpADD(value int) {
	ula.Ac.Value += value
}

func (ula *ULA) ExecuteOpSUB(value int) {
	ula.Ac.Value -= value
}

func (ula *ULA) ExecuteOpMUL(value int) {
	ula.Ac.Value *= value
}

func (ula *ULA) ExecuteOpDIV(value int) {
	ula.Ac.Value /= value
}

func (ula *ULA) ExecuteOpAND(value int) {
	ula.Ac.Value = int(int8(ula.Ac.Value) & int8(value))
}

func (ula *ULA) ExecuteOpOR(value int) {
	ula.Ac.Value = int(int8(ula.Ac.Value) | int8(value))
}

func (ula *ULA) ExecuteOpXOR(value int) {
	ula.Ac.Value = int(int8(ula.Ac.Value) ^ int8(value))
}
