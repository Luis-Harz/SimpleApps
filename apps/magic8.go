package apps

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Main9() {
	Clear()
	welcometext := "----Welcome to Magic 8-Ball----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	for {
		var input string
		fmt.Println("[Type 'exit' to exit]")
		fmt.Println("Ask your question:")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		rn := rand.Intn(2)
		Clear()
		if rn == 0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
