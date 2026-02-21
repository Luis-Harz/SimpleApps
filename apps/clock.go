package apps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Main8() {
	Clear()
	welcometext := "----Welcome to Clock----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	time.Sleep(time.Second)
	exitChan := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				exitChan <- true
				return
			}
		}
	}()
	for {
		select {
		case <-exitChan:
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			return
		default:
			Clear()
			now := time.Now()
			Time := now.Format("15:04:05")
			fmt.Println("[Press 'q' or type 'exit' to exit]")
			printNumber(Time)
			fmt.Printf("%s ", ConfigData.Prompt)
			time.Sleep(time.Second)
		}
	}
}
