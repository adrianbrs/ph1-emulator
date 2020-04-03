package ula

import (
	"ph1-emulator/regs"
)

type ULA struct{}

var ula = &ULA{}

func GetInstance() *ULA {
	return ula
}

func (ula *ULA) ExecuteOpNOT() {
	regs.RegAC.Value = int(^int8(regs.RegAC.Value))
}

func (ula *ULA) ExecuteOpADD(value int) {
	regs.RegAC.Value += value
}

func (ula *ULA) ExecuteOpSUB(value int) {
	regs.RegAC.Value -= value
}

func (ula *ULA) ExecuteOpMUL(value int) {
	regs.RegAC.Value *= value
}

func (ula *ULA) ExecuteOpDIV(value int) {
	regs.RegAC.Value /= value
}

func (ula *ULA) ExecuteOpAND(value int) {
	regs.RegAC.Value = int(int8(regs.RegAC.Value) & int8(value))
}

func (ula *ULA) ExecuteOpOR(value int) {
	regs.RegAC.Value = int(int8(regs.RegAC.Value) | int8(value))
}

func (ula *ULA) ExecuteOpXOR(value int) {
	regs.RegAC.Value = int(int8(regs.RegAC.Value) ^ int8(value))
}
