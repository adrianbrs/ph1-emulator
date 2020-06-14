package numbers

import (
	"fmt"
	config "ph1-emulator/constants"
	"strconv"
)

// ValidateUint8 valida se um inteiro cabe em um inteiro de 8 bits
// Retorna um inteiro de 8 bits sem sinal
func ValidateUint8(val int) (res uint8, err error) {
	num, err := strconv.ParseUint(fmt.Sprintf("%d", val), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(num), nil
}

// ValidateInt8 valida se um inteiro cabe em um inteiro de 8 bits
// Retorna um inteiro de 8 bits com sinal
func ValidateInt8(val int) (res int8, err error) {
	num, err := ValidateUint8(val)
	if err != nil {
		return 0, err
	}
	return int8(num), nil
}

// ValidateAddress valida o tamanho de um endereço de memória
func ValidateAddress(addr int) (int, error) {
	res, err := strconv.ParseUint(fmt.Sprintf("%d", addr), 10, config.AddrLength)
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

// ValidateWord valida o tamanho de uma palavra de memória
func ValidateWord(val int) (int, error) {
	resUint, errUint := strconv.ParseUint(fmt.Sprintf("%d", val), 10, config.WordLength)

	if errUint != nil {
		resInt, errInt := strconv.ParseInt(fmt.Sprintf("%d", val), 10, config.WordLength)
		if errInt != nil {
			return 0, errInt
		}
		return int(resInt), nil
	}
	return int(resUint), nil
}
