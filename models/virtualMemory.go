package models

import (
	"log"
	"strconv"
)

//VirtualMemory representation of a virtualMemory structure
type VirtualMemory struct {
	Address, Data string
}

const (
	hexadecimalBaseValue = 16
	bitsLimit            = 8
)

//ValidateAddress validate if requested address is bigger than 8 bits
func ValidateAddress(request string) {
	_, err := strconv.ParseInt(request, hexadecimalBaseValue, bitsLimit)

	if err != nil {
		log.Panic(map[string]string{
			"message:": "Request address is bigger than 8 bits"})
	}
}
