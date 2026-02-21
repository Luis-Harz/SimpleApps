package apps

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/term"
)

var chatMu sync.Mutex
var currentInput string

func reprintInputLine() {
	fmt.Printf("\r\033[K%s %s", ConfigData.Prompt, currentInput)
}

func printChatMessage(message string) {
	chatMu.Lock()
	defer chatMu.Unlock()
	fmt.Printf("\r\033[K%s\n", message)
	reprintInputLine()
}

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

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error entering raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Connected to ChatServer", SERVER)
	fmt.Println("[Type 'exit' or 'q' to exit, Ctrl+C to force quit]\r")

	done := make(chan bool, 1)
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				printChatMessage("Disconnected from server.")
				done <- true
				return
			}
			printChatMessage(string(message))
		}
	}()

	chatMu.Lock()
	currentInput = ""
	reprintInputLine()
	chatMu.Unlock()

	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		b := buf[0]
		if b == 3 || b == 4 {
			break
		}
		chatMu.Lock()
		if b == 13 {
			text := strings.TrimSpace(currentInput)
			currentInput = ""
			fmt.Print("\r\n")
			chatMu.Unlock()

			if text == "" {
				chatMu.Lock()
				reprintInputLine()
				chatMu.Unlock()
				continue
			}

			conn.WriteMessage(websocket.TextMessage, []byte(text))

			if strings.ToLower(text) == "exit" || strings.ToLower(text) == "q" {
				break
			}

			chatMu.Lock()
			reprintInputLine()
			chatMu.Unlock()

		} else if b == 127 || b == 8 {
			if len(currentInput) > 0 {
				currentInput = currentInput[:len(currentInput)-1]
			}
			reprintInputLine()
			chatMu.Unlock()

		} else if b >= 32 {
			currentInput += string(b)
			reprintInputLine()
			chatMu.Unlock()

		} else {
			chatMu.Unlock()
		}
	}
	select {
	case <-done:
	default:
	}

	fmt.Print("\r\nClient Closed.\r\n")
}
