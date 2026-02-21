package apps

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Main6() {
	scanner := bufio.NewScanner(os.Stdin)
	welcometext := "----Welcome to Countdown----"
	minuses := (len(welcometext)/2 - len(" Bye! ")/2)

	for {
		Clear()
		fmt.Println(welcometext)

		fmt.Println("How many Seconds (or type 'exit' to exit)? ")
		fmt.Printf("%s ", ConfigData.Prompt)
		if !scanner.Scan() {
			break // EOF
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(2 * time.Second)
			Clear()
			break
		}

		seconds, err := strconv.Atoi(input)
		if err != nil || seconds < 0 {
			Clear()
			fmt.Println("Please give a valid input!")
			time.Sleep(2 * time.Second)
			continue
		}

		for seconds > 0 {
			Clear()
			h := seconds / 3600
			m := (seconds % 3600) / 60
			s := seconds % 60
			fmt.Printf("%02d:%02d:%02d\n", h, m, s)
			seconds--
			time.Sleep(time.Second)
		}

		Clear()
		fmt.Println("00:00:00")
		fmt.Println("Time's Over!")
		time.Sleep(2 * time.Second)
	}
}
