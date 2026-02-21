package apps

import (
	"fmt"
	"strings"
	"time"
)

func Main7() {
	welcometext := "----Welcome to Timer----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var start time.Time
	var paused time.Duration
	running := false

	for {
		Clear()
		fmt.Println(welcometext)
		fmt.Println("Options:")
		fmt.Println("[start, stop, reset, show, exit]")

		var cmd string
		fmt.Println("Enter command: ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&cmd)

		switch cmd {
		case "start":
			if !running {
				start = time.Now().Add(-paused)
				running = true
			}
		case "stop":
			if running {
				paused = time.Since(start)
				running = false
			}
		case "reset":
			paused = 0
			running = false
		case "show":
			if running {
				fmt.Println("Timer:", formatDuration(time.Since(start)))
			} else {
				fmt.Println("Timer:", formatDuration(paused))
			}
			time.Sleep(2 * time.Second)
		case "exit":
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			return
		default:
			fmt.Println("Unknown command!")
			time.Sleep(2 * time.Second)
		}
	}
}
