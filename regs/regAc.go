package regs

import (
	"log"
	c "ph1-emulator/constants"
	"ph1-emulator/numbers"
)

// RegAC = Registrador Program Counter
type regAC struct {
	value int
}

// SetValue atualiza o valor do registrador
func (ac *regAC) SetValue(value int) {
	// Necessário verificar overflow antes de atualizar o valor
	val, err := numbers.ValidateWord(value)

	// Overflow
	if err != nil {
		log.Fatalf(c.AcRangeOverflow, value)
	}

	ac.value = val
}

// GetValue retorna o valor atual (puro) do registrador
func (ac *regAC) GetValue() int {
	return ac.value
}

// GetWordValue retorna o valor do registrador devidamente ajustado
// para 8 bits
func (ac *regAC) GetWordValue() int {
	return int(int8(ac.GetValue()))
}

// RegAC => Registrador Acumulador
var RegAC = &regAC{value: 0}
