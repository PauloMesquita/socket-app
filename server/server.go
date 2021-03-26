package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"trabalho_webmotors/webmotors"
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
		clients: make(map[net.Conn]string),
		registerUser: make(chan ConnectionName),
		unregisterUser: make(chan net.Conn),
		messages: make(chan ConnectionName),
		webmotors: webmotors.NewWebMotos(),
	}
	
	go server.handleEvents() // Lidar com eventos que chegam pelos canais

	// Lidar com novas conexões
	for {
		socket, err := listener.Accept() // Aceita a conexão com o cliente
		if err != nil {
			fmt.Println("Erro ao conectar com cliente: ", err.Error())
		}
		socket.Write([]byte("Digite seu nome: "))
		go server.receive(socket) // Service for receiving messages
	}

}

func getPorta() (string, error) {

	if len(os.Args) == 1 {
		return "", errors.New("é necessário enviar a porta")
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

type ConnectionName struct {
	connection net.Conn
	name string
}

type Server struct {
	clients map[net.Conn]string
	registerUser chan ConnectionName
	unregisterUser chan net.Conn
	messages chan ConnectionName
	webmotors *webmotors.Webmotors
}

func (s *Server) receive(client net.Conn) {
	receiveMessage := make([]byte, 512)
	

	for {
		clearArray(receiveMessage)
		messageSize, err := client.Read(receiveMessage)
		if err != nil {
			// Client disconnects
			s.unregisterRegisteredUser(client)
			break
		}
		if messageSize == 0 {
			continue
		}
		message := string(receiveMessage[:messageSize-1])
		s.messages <- ConnectionName{client, message}
	}
}

func (s *Server) handleEvents() {
	for cn := range s.messages {
		s.registerUnregisteredUser(&cn)
		fmt.Printf("%s\n", cn.name)
	}
}

func (s *Server) registerUnregisteredUser(cn *ConnectionName) {
	// If client not connected
	if _, exist := s.clients[cn.connection]; !exist {
		// Verify if name is already taken
		taken := false
		for _, v := range s.clients {
			fmt.Printf("%s %s", cn.name, v)
			if cn.name == v  {
				taken = true
			} 
		}
		// Send registerUser message
		if !taken {
			s.clients[cn.connection] = cn.name
			fmt.Printf("O cliente %s foi cadastrado\n", cn.name)
		}else{
			cn.connection.Write([]byte("Nome já utilizado. Digite seu nome: "))
		}
		
	}
}

func (s *Server) unregisterRegisteredUser(c net.Conn) {
	value, exists := s.clients[c]
	if exists {
		delete(s.clients, c)
		c.Close()
		fmt.Printf("O cliente %s foi descadastrado\n", value)
	}
}

// ### Server ###


