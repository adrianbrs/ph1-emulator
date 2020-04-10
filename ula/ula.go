package ula

import (
	"ph1-emulator/memory"
	"ph1-emulator/regs"
)

//ULA struct
type ULA struct{}

var ula = &ULA{}

//GetInstance create a ULA instance
func GetInstance() *ULA {
	return ula
}

//ExecuteOpNOT execute bitwise NOT operation
func (ula *ULA) ExecuteOpNOT() {
	regs.RegAC.SetValue(int(^int8(regs.RegAC.GetValue())))
}

//ExecuteOpADD execute sum operation
func (ula *ULA) ExecuteOpADD(addr int) {
	regs.RegAC.SetValue(regs.RegAC.GetWordValue() + memory.VirtualMemory.GetWordValue(addr))
}

//ExecuteOpSUB execute subtraction operation
func (ula *ULA) ExecuteOpSUB(addr int) {
	regs.RegAC.SetValue(regs.RegAC.GetWordValue() - memory.VirtualMemory.GetWordValue(addr))
}

//ExecuteOpMUL execute multiplication operation
func (ula *ULA) ExecuteOpMUL(addr int) {
	regs.RegAC.SetValue(regs.RegAC.GetWordValue() * memory.VirtualMemory.GetWordValue(addr))
}

//ExecuteOpDIV execute division operation
func (ula *ULA) ExecuteOpDIV(addr int) {
	regs.RegAC.SetValue(regs.RegAC.GetWordValue() / memory.VirtualMemory.GetWordValue(addr))
}

//ExecuteOpAND execute bitwise AND operation
func (ula *ULA) ExecuteOpAND(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.SetValue(int(int8(regs.RegAC.GetValue()) & int8(value)))
}

//ExecuteOpOR execute bitwise OR operation
func (ula *ULA) ExecuteOpOR(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.SetValue(int(int8(regs.RegAC.GetValue()) | int8(value)))
}

//ExecuteOpXOR execute bitwise XOR operation
func (ula *ULA) ExecuteOpXOR(addr int) {
	value := memory.VirtualMemory.GetValue(addr)
	regs.RegAC.SetValue(int(int8(regs.RegAC.GetValue()) ^ int8(value)))
}
