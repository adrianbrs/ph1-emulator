package regs

type regPC struct {
	value int
}

// Increment adiciona 1 ao valor do registrador
func (reg *regPC) Increment() {
	reg.value++
}

// GetValue retorna o valor do registrador
func (reg *regPC) GetValue() int {
	return reg.value
}

// SetValue altera o valor do registrador
func (reg *regPC) SetValue(val int) {
	reg.value = val
}

// RegPC => Registtrador Program Counter
var RegPC = &regPC{value: 0}
