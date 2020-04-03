package regs

// RegAC = Registrador Program Counter
type regAC struct {
	Value int
}

// RegAC => Registrador Acumulador
var RegAC = &regAC{Value: 0}
