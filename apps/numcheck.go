package apps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Main1() {
	Clear()
	welcometext := "----Welcome to NumberChecker----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	for {
		var input string
		fmt.Println("Type a Number(or 'exit' to exit): ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please Type a valid Number!")
			continue
		}
		Clear()
		if number%2 == 0 {
			fmt.Printf("Number %d is even\n", number)
		} else {
			fmt.Printf("Number %d is odd\n", number)
		}
	}
}
