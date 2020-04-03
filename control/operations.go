package control

import (
	"ph1-emulator/memory"
	"ph1-emulator/regs"
)

func (uc *unityControl) ExecuteOpNOP() {}

func (uc *unityControl) ExecuteOpLDR(value int) {
	regs.RegAC.Value = memory.VirtualMemory.GetValue(value)
}

func (uc *unityControl) ExecuteOpSTR(value int) {
	memory.VirtualMemory.SetValue(value, regs.RegAC.Value)
}

func (uc *unityControl) ExecuteOpJMP(value int) {
	regs.RegPC.SetValue(value)
}

func (uc *unityControl) ExecuteOpJEQ(value int) {
	if regs.RegAC.Value == 0 {
		regs.RegPC.SetValue(value)
	}
}

func (uc *unityControl) ExecuteOpJG(value int) {
	if regs.RegAC.Value > 0 {
		regs.RegPC.SetValue(value)
	}
}

func (uc *unityControl) ExecuteOpJL(value int) {
	if regs.RegAC.Value < 0 {
		regs.RegPC.SetValue(value)
	}
}

func (uc *unityControl) ExecuteOpHLT() {
	uc.Houte = true
}
