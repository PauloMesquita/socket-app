package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main(){
	port, err := getPorta()
	if err != nil {
		exitProgramAndPrintMessage(err.Error())
	}
	listener, err := net.Listen("tcp",":"+port)
	if err != nil {
		exitProgramAndPrintMessage("Error ao ouvir porta "+port )
	}
	defer listener.Close()

	server := Server{
		clients: make(map[net.Conn]bool),
		registerUser: make(chan net.Conn),
		unregisterUser: make(chan net.Conn),
	}
	
	go server.handleEvents() // Lidar com eventos que chegam pelos canais

	// Lidar com novas conexões
	for {
		socket, err := listener.Accept() // Aceita a conexão com o cliente
		if err != nil {
			fmt.Println("Erro ao conectar com cliente: ", err.Error())
		}
		server.registerUser <- socket
		go server.receive(socket) // Receber mensagens do cliente que acabou de conectar
	}

}

func getPorta() (string, error) {

	if len(os.Args) == 1 {
		return "", errors.New("É necessário enviar a porta")
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

//  ### Server ###
type Server struct {
	clients map[net.Conn]bool
	registerUser chan net.Conn
	unregisterUser chan net.Conn
}

func (s *Server) receive(client net.Conn) {
	receiveMessage := make([]byte, 512)

	for {
		clearArray(receiveMessage)
		messageSize, err := client.Read(receiveMessage)
		if err != nil {
			// Client disconnects
			s.unregisterUser <- client
			client.Close()
			break
		}
		if messageSize == 0 {
			continue
		}
		fmt.Println(string(receiveMessage[:messageSize-1]))
		//client.Write([]byte("Mensagem recebida "+ string(receiveMessage[:messageSize])))
	}
}

func (s *Server) handleEvents() {
	for {
		select {
			case socket := <- s.registerUser:
				s.clients[socket] = true
				fmt.Println("Um novo cliente foi cadastrado")
			case socket := <- s.unregisterUser:
				_, exists := s.clients[socket]
				if exists {
					delete(s.clients, socket)
					fmt.Println("Um cliente foi descadastrado")
				}
		}
	}
}
// ### Server ###


