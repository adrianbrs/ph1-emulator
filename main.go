package main

import (
	"fmt"
	"os"
	"os/signal"
	"ph1-emulator/control"
	"ph1-emulator/input"
	"ph1-emulator/logger"
	"syscall"
)

func main() {
	// Inicia a rotina de espera para o sinal de interrupção
	halt := make(chan bool, 1)
	done := make(chan bool, 1)
	SetupExitHandler(halt)

	// Cria uma goroutine para ficar aguardando o sinal de parada
	go func() {
		// Aguarda pelo sinal de parada (HLT)
		success := <-halt

		// Exibe as informações do estado final dos registradores e memória
		logger.LogFinalState(control.UnityControl.Counter)

		// Se o valor recebido em halt for falso, quer dizer que o
		// programa foi interrompido pelo usuário (Ctrl+C)
		if !success {
			// Exibe mensagem de parada
			fmt.Print("\n\n")
			fmt.Println("\r- Ctrl+C pressed in terminal, program exited.")
		}

		// Envia true para o canal done para que saia da função main
		// e termine a execução do programa
		done <- true
		os.Exit(0)
	}()

	// Lê o arquivo de entrada contendo as instruções
	input.RequestInput()

	// Inicia a unidade central de processamento
	control.UnityControl.Start(halt)

	// Aguarda o sinal de parada para terminar a execução,
	// assim garantimos que a goroutine acima será executada
	// exibindo o log final
	<-done
}

// SetupExitHandler inicia uma goroutine que ficara esperando
// um sinal de interrupção para exibir o estado final da execução
// antes dela ser interrompida
func SetupExitHandler(halt chan bool) {
	stopChannel := make(chan os.Signal)

	signal.Notify(stopChannel, os.Interrupt, syscall.SIGTERM)

	// Cria uma goroutine em paralelo para ficar aguardando o sinal
	// de interrupção
	go func() {
		// Aguarda pelo sinal de interrupção do terminal
		<-stopChannel

		// Interrompe a execução da unidade de controle
		control.UnityControl.Houte = true

		// Envia o sinal de halt (HLT) caso a unidade de controle ainda
		// não tenha sido iniciada para enviar o sinal
		halt <- false
	}()
}
