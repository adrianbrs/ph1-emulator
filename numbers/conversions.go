package numbers

import (
	"fmt"
	"strconv"
)

// HexToUint converte hex para inteiro sem sinal de 64 bits
func HexToUint(str string, bits int) (res uint64, err error) {
	res, err = strconv.ParseUint(str, 16, bits)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// IntToHex converte inteiro para hexadecimal (string)
func IntToHex(value int, digits int) string {
	return fmt.Sprintf(("%0" + fmt.Sprintf("%d", digits) + "X"), value)
}
