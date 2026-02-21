package apps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Help Functions main20
func printenabledletters(enabled []string, names []string) {
	lines := len(enabled)
	var itemenabled string
	for i := 0; i < lines; i++ {
		if enabled[i] == "1" {
			itemenabled = "X"
		} else {
			itemenabled = " "
		}
		fmt.Println("[" + strconv.Itoa(i) + "]" + "[" + itemenabled + "] " + names[i])
	}
}

func setvariablesletters(numbersenabled bool, lettersupperenabled bool, letterslowerenabled bool, lettersspecialenabled bool) []string {
	enabled := make([]string, 4)
	if numbersenabled == true {
		enabled[0] = "1"
	} else if numbersenabled == false {
		enabled[0] = "0"
	}
	if lettersupperenabled == true {
		enabled[1] = "1"
	} else if lettersupperenabled == false {
		enabled[1] = "0"
	}
	if letterslowerenabled == true {
		enabled[2] = "1"
	} else if letterslowerenabled == false {
		enabled[2] = "0"
	}
	if lettersspecialenabled == true {
		enabled[3] = "1"
	} else if lettersspecialenabled == false {
		enabled[3] = "0"
	}
	return enabled
}
func combinelists(enabled []string, list1, list2, list3, list4 []string) []string {
	allLists := [][]string{list1, list2, list3, list4}
	var listscombined []string
	for i := 0; i < len(enabled) && i < len(allLists); i++ {
		if enabled[i] == "1" {
			listscombined = append(listscombined, allLists[i]...)
		}
	}
	return listscombined
}

//END Help Functions main20

func Main20() {
	welcometext := "----Welcome to PasswordGen----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	names := []string{"Numbers", "Upper Letters", "Lower Letters", "Special Letters"}
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	lettersupper := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	letterslower := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	lettersspecial := []string{"!", "?"}
	enabled0 := true
	enabled1 := true
	enabled2 := true
	enabled3 := false
	enabled := setvariablesletters(enabled0, enabled1, enabled2, enabled3)
	for {
		Clear()
		fmt.Println(welcometext)
		printenabledletters(enabled, names)
		fmt.Println("[toggle by using numbers or type 'gen' to generate or type 'exit' to exit]")
		fmt.Print(ConfigData.Prompt + " ")
		input := ReadInput()
		if input == "0" {
			enabled0 = !enabled0
			enabled = setvariablesletters(enabled0, enabled1, enabled2, enabled3)
		} else if input == "1" {
			enabled1 = !enabled1
			enabled = setvariablesletters(enabled0, enabled1, enabled2, enabled3)
		} else if input == "2" {
			enabled2 = !enabled2
			enabled = setvariablesletters(enabled0, enabled1, enabled2, enabled3)
		} else if input == "3" {
			enabled3 = !enabled3
			enabled = setvariablesletters(enabled0, enabled1, enabled2, enabled3)
		} else if input == "gen" {
			combinedlists := combinelists(enabled, numbers, lettersupper, letterslower, lettersspecial)
			fmt.Println("[Type the length]")
			fmt.Print(ConfigData.Prompt + " ")
			input2 := ReadInput()
			input2int, err := strconv.Atoi(input2)
			if err != nil {
				panic(err)
			}
			var allletters string
			for i := 0; i < input2int; i++ {
				letter := combinedlists[Random(0, len(combinedlists)-1)]
				allletters += letter
			}
			fmt.Println("Password: " + allletters)
			time.Sleep(time.Second * 5)
		} else if input == "help" {
			fmt.Println("---Commands---\n 0\n 1\n 2\n 3\n gen\n exit\n help")
			time.Sleep(time.Second * 2)
		} else if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		} else {
			fmt.Println("Command Not Found. Type 'help' to see all commands")
			time.Sleep(time.Second * 2)
		}
	}
}
