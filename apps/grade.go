package apps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Main2() {
	Clear()
	welcometext := "----Welcome to GradeChecker----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	for {
		var maxscorestring string
		var scoreStr string
		fmt.Println("What's the Maximal Score(type 'exit' to exit)? ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&maxscorestring)
		if maxscorestring == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		maxscorestring = strings.TrimSpace(maxscorestring)
		maxscorestring = strings.ReplaceAll(maxscorestring, ",", ".")
		fmt.Println("What's the Student's Score? ")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&scoreStr)
		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			fmt.Println("Please type a valid number!")
			continue
		}
		maxscoreFloat, err := strconv.ParseFloat(maxscorestring, 64)
		if err != nil {
			fmt.Println("Invalid max score!")
			continue
		}
		percent := score / maxscoreFloat * 100
		Clear()
		switch {
		case percent >= 90:
			fmt.Printf("With a Score of %.1f the Student get's an A\n", score)
		case percent >= 80:
			fmt.Printf("With a Score of %.1f the Student get's a B\n", score)
		case percent >= 70:
			fmt.Printf("With a Score of %.1f the Student get's a C\n", score)
		case percent >= 60:
			fmt.Printf("With a Score of %.1f the Student get's a D\n", score)
		default:
			fmt.Printf("With a Score of %.1f the Student get's an F\n", score)
		}
	}
}
