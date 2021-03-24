package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)


func main() {
	inputSocket, err := getSocket()
	if err != nil {
		exitProgramAndPrintMessage(err.Error())
	}
	socket, err := net.Dial("tcp", inputSocket)
	if err != nil {
		exitProgramAndPrintMessage("Erro ao tentar conectar com servidor: "+err.Error())
	}
	fmt.Println("Conexão realizada com sucesso")
	
	// Receber mensagens
	go receiveMessage(socket)

	// Enviar mensagens
	reader := bufio.NewReader(os.Stdin)
	for {
		mensagem, err := reader.ReadString('\n')
		if err == io.EOF {
			return // Cliente fechou a conexão
		}
		socket.Write([]byte(mensagem))
	}
}

func receiveMessage(server net.Conn) {
	message := make([]byte, 512)

	for {
		clearArray(message)
		messageSize, err := server.Read(message)
		if err != nil {
			fmt.Println("Erro ao ler mensagem: ", err)
			server.Close()
			exitProgramAndPrintMessage("Servidor fechou")
		}
		if messageSize == 0 {
			continue
		}
		fmt.Println(string(message[:messageSize]))
	}
}

func getSocket() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("É necessário enviar o socket")
	}
	return os.Args[1], nil
}

func exitProgramAndPrintMessage(errormessage string) {
	fmt.Println(errormessage)
	os.Exit(1)
}

func clearArray(array []byte) {
	for i:=0; i < len(array); i++{
		array[i] = 0
	}
}