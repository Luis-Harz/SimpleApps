package apps

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func Main4() {
	Clear()
	welcometext := "----Welcome to N2B----"
	fmt.Println(welcometext)
	fmt.Println("------Number2Bar------")
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var inputstr string
	var maximalstr string
	for {
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Error getting Terminal size:", err)
			return
		}
		widthstr := strconv.Itoa(width - 2)
		fmt.Println("Please set the maximum(Terminal width: " + widthstr + ")(or type 'exit' to exit): ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&maximalstr)
		if maximalstr == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		maximal, err := strconv.Atoi(maximalstr)
		if err != nil || maximal > width {
			Clear()
			fmt.Println("Please give a valid number!")
			continue
		}
		fmt.Println("Please give a number(max." + maximalstr + "): ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&inputstr)
		if inputstr == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		inputfloat, err := strconv.ParseFloat(inputstr, 64)
		if err != nil {
			Clear()
			fmt.Println("Please type a valid number!")
			continue
		}
		input := int(inputfloat)
		if input < 0 || input > maximal || input > width-2 {
			Clear()
			fmt.Println("Please type a valid number!")
			continue
		}
		max := maximal
		parts := max - input
		Clear()
		fmt.Println("[" + strings.Repeat("#", input) + strings.Repeat(" ", parts) + "]")
	}
}
