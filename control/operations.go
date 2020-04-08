package control

import (
	"ph1-emulator/memory"
	"ph1-emulator/regs"
)

//ExecuteOpNOP dont execute nothing
func (uc *unityControl) ExecuteOpNOP() {}

//ExecuteOpLDR execute a load operation
func (uc *unityControl) ExecuteOpLDR(value int) {
	regs.RegAC.Value = memory.VirtualMemory.GetValue(value)
}

//ExecuteOpSTR execute a store operation
func (uc *unityControl) ExecuteOpSTR(value int) {
	memory.VirtualMemory.SetValue(value, regs.RegAC.Value)
}

//ExecuteOpJMP execute a jump operation
func (uc *unityControl) ExecuteOpJMP(value int) {
	regs.RegPC.SetValue(value)
}

//ExecuteOpJEQ execute a conditional equals operation
func (uc *unityControl) ExecuteOpJEQ(value int) {
	if regs.RegAC.Value == 0 {
		regs.RegPC.SetValue(value)
	}
}

//ExecuteOpJG execute a conditional bigger operation
func (uc *unityControl) ExecuteOpJG(value int) {
	if regs.RegAC.Value > 0 {
		regs.RegPC.SetValue(value)
	}
}

//ExecuteOpJL execute a conditional lower operation
func (uc *unityControl) ExecuteOpJL(value int) {
	if regs.RegAC.Value < 0 {
		regs.RegPC.SetValue(value)
	}
}

//ExecuteOpHLT execute a halt operation
func (uc *unityControl) ExecuteOpHLT() {
	uc.Houte = true
}
