package registrators

// RegAC = Registrador Program Counter
type RegAC struct {
	value int
}

// GetValue retorna o valor do registrador
func (reg *RegAC) GetValue() int {
	return reg.value
}

// SetValue altera o valor do registrador
func (reg *RegAC) SetValue(val int) *RegAC {
	reg.value = val
	return reg
}

// NewRegAC cria uma nova instância do registrador AC
func NewRegAC() (pc *RegAC) {
	pc = &RegAC{value: 0}
	return pc
}
