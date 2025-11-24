package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Connected to Chatroom. Type 'exit' to go.\n")

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chatroom, Bye :)")
			break
		}

		var history []string

		err := client.Call("ChatService.SendMessage", text, &history)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Println("---- Chat History ----")
		for _, msg := range history {
			fmt.Print(msg)
		}
		fmt.Println("----------------------")

	}
}
