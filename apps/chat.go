package apps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Main19() {
	var SERVER string
	fmt.Println("Enter Server URL [default: wss://chat.bolucraft.uk/chat]")
	fmt.Print(ConfigData.Prompt + " ")
	fmt.Scanln(&SERVER)
	if SERVER == "" {
		SERVER = "wss://chat.bolucraft.uk/chat"
	}

	conn, _, err := websocket.DefaultDialer.Dial(SERVER, nil)
	if err != nil {
		fmt.Println("Error Connecting:", err)
		time.Sleep(time.Second * 2)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to ChatServer", SERVER)
	fmt.Println("[Type 'q' or 'exit' to exit]")

	done := make(chan bool)
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("\nDisconnected from server.")
				done <- true
				return
			}
			fmt.Println("\n" + string(message))
			fmt.Print(ConfigData.Prompt + " ")
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(ConfigData.Prompt + " ")
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		conn.WriteMessage(websocket.TextMessage, []byte(text))
		if strings.ToLower(text) == "exit" || strings.ToLower(text) == "q" {
			break
		}
	}

	<-done
	fmt.Println("Client Closed.")
}
