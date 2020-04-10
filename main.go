package main

import (
	"ph1-emulator/control"
	"ph1-emulator/input"
	"ph1-emulator/logger"
)

func main() {
	// Lê o arquivo de entrada contendo as instruções
	input.ReadInstructionsFile("Input file: ")

	// Inicia a unidade central de processamento
	control.UnityControl.Start()

	// Exibe as informações do estado final dos registradores e memória
	logger.LogFinalState()
}
