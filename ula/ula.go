package ula

import (
	"ph1-emulator/memory"
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

func (ula *ULA) ExecuteOpADD(addr int) {
	regs.RegAC.Value += memory.VirtualMemory.GetValue(addr)
}

func (ula *ULA) ExecuteOpSUB(addr int) {
	regs.RegAC.Value -= memory.VirtualMemory.GetValue(addr)
}

func (ula *ULA) ExecuteOpMUL(addr int) {
	regs.RegAC.Value *= memory.VirtualMemory.GetValue(addr)
}

func (ula *ULA) ExecuteOpDIV(addr int) {
	regs.RegAC.Value /= memory.VirtualMemory.GetValue(addr)
}

func (ula *ULA) ExecuteOpAND(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.Value = int(int8(regs.RegAC.Value) & int8(value))
}

func (ula *ULA) ExecuteOpOR(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.Value = int(int8(regs.RegAC.Value) | int8(value))
}

func (ula *ULA) ExecuteOpXOR(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.Value = int(int8(regs.RegAC.Value) ^ int8(value))
}
