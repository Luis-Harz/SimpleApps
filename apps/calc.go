package apps

import (
	"fmt"
	"strings"
	"time"
)

func Main11() {
	Clear()
	welcometext := "----Welcome to Calculator----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var input string
	for {
		fmt.Println("[Type 'exit' to exit]")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		tokens := tokenize(input)
		parser := Parser{tokens: tokens}
		result := parser.expr()
		fmt.Println("=", result)
		time.Sleep(time.Second * 2)
		Clear()
	}
}
