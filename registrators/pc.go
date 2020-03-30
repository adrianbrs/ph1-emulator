package registrators

// RegPC = Registrador Program Counter
type RegPC struct {
	value int
}

// Increase adiciona 1 ao valor do registrador
func (reg *RegPC) Increase() *RegPC {
	reg.value++
	return reg
}

// GetValue retorna o valor do registrador
func (reg *RegPC) GetValue() int {
	return reg.value
}

// SetValue altera o valor do registrador
func (reg *RegPC) SetValue(val int) *RegPC {
	reg.value = val
	return reg
}

// NewRegPC cria uma nova instância do registrador PC
func NewRegPC() (pc *RegPC) {
	pc = &RegPC{value: 0}
	return pc
}
