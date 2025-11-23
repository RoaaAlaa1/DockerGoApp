package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ChatService struct {
	messages []string
}

func (c *ChatService) SendMessage(msg string, reply *[]string) error {
	msg = msg + "\n"
	c.messages = append(c.messages, msg)
	*reply = c.messages
	fmt.Printf("Received: %s\n", msg)
	return nil
}

func main() {
	server := new(ChatService)
	rpc.Register(server)

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server running on port 1234")

	if err != nil {
		fmt.Println("Error registering service:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
