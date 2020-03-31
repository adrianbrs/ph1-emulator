package numbers

import (
	"fmt"
	"log"
	"strconv"
)

// HexToInt converte hex para inteiro sem sinal de 64 bits
func HexToInt(str string, bits int) int {
	res, err := strconv.ParseUint(str, 16, bits)
	if err != nil {
		log.Fatalf("Could not convert hex '%s' to int.", str)
	}
	return int(res)
}

// IntToHex converte inteiro para hexadecimal (string)
func IntToHex(value int, digits int) string {
	return fmt.Sprintf(("%0" + fmt.Sprintf("%d", digits) + "X"), value)
}
