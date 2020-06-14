package numbers

import (
	"fmt"
	"log"
	config "ph1-emulator/constants"
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
func IntToHex(value int) string {
	return fmt.Sprintf(("%0" + fmt.Sprintf("%d", config.HexDigits) + "X"), uint8(value))
}
