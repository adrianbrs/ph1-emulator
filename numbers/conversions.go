package numbers

import (
	"fmt"
	"strconv"
)

// HexToInt converte hex para inteiro sem sinal de 64 bits
func HexToInt(str string, bits int) (int, error) {
	res, err := strconv.ParseUint(str, 16, bits)
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

// IntToHex converte inteiro para hexadecimal (string)
func IntToHex(value int, digits int) string {
	return fmt.Sprintf(("%0" + fmt.Sprintf("%d", digits) + "X"), value)
}
