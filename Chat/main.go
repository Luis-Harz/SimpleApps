package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

const MaxMessageLength = 200

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	clients := make(map[*websocket.Conn]string)
	mutex := &sync.Mutex{}

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}

		clientIP := r.Header.Get("CF-Connecting-IP")
		if clientIP == "" {
			clientIP = conn.RemoteAddr().String()
		}

		mutex.Lock()
		clients[conn] = clientIP
		mutex.Unlock()

		fmt.Println("New Client connected:", clientIP)

		go handleClient(conn, clients, mutex)
	})

	fmt.Println("Chatserver running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleClient(conn *websocket.Conn, clients map[*websocket.Conn]string, mutex *sync.Mutex) {
	defer func() {
		mutex.Lock()
		clientIP := clients[conn]
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", clientIP)
	}()

	clientIP := clients[conn]
	conn.WriteMessage(websocket.TextMessage, []byte("Welcome to the Chat! Type 'exit' or 'q' to exit."))

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			return
		}

		msg := strings.TrimSpace(string(msgBytes))
		if msg == "" {
			continue
		}
		if strings.ToLower(msg) == "exit" || strings.ToLower(msg) == "q" {
			conn.WriteMessage(websocket.TextMessage, []byte("Bye!"))
			return
		}
		if len(msg) > MaxMessageLength {
			if MaxMessageLength > 3 {
				msg = msg[:MaxMessageLength-3] + "..."
			} else {
				msg = msg[:MaxMessageLength]
			}
		}

		mutex.Lock()
		for c := range clients {
			if c != conn {
				c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", clientIP, msg)))
			}
		}
		mutex.Unlock()
	}
}
