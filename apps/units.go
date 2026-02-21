package apps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Main3() {
	Clear()
	welcometext := "----Welcome to UnitConverter----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)

	for {
		// Length
		fmt.Println(" [0]  cm => m")
		fmt.Println(" [1]  m => cm")
		fmt.Println(" [2]  cm => dm")
		fmt.Println(" [3]  dm => cm")
		fmt.Println(" [4]  m => dm")
		fmt.Println(" [5]  dm => m")
		// Time
		fmt.Println(" [6]  seconds => minutes")
		fmt.Println(" [7]  seconds => hours")
		fmt.Println(" [8]  minutes => seconds")
		fmt.Println(" [9]  hours => seconds")
		fmt.Println(" [10] minutes => hours")
		fmt.Println(" [11] hours => minutes")
		// Mass
		fmt.Println(" [12] g => kg")
		fmt.Println(" [13] kg => g")
		// Volume
		fmt.Println(" [14] ml => l")
		fmt.Println(" [15] l => ml")

		var choice string
		fmt.Println("What Conversion (or type 'exit' to exit)?")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}

		choiceint, err := strconv.Atoi(choice)
		if err != nil {
			Clear()
			fmt.Println("Please type a valid Number!")
			continue
		}

		var number string
		fmt.Println("How much?")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&number)
		numberint, err := strconv.ParseFloat(number, 64)
		if err != nil {
			Clear()
			fmt.Println("Please type a valid Number!")
			continue
		}

		var outputint float64
		var unit string

		if choiceint == 0 {
			outputint = numberint * 0.01
			unit = "m"
		}
		if choiceint == 1 {
			outputint = numberint * 100
			unit = "cm"
		}
		if choiceint == 2 {
			outputint = numberint * 0.1
			unit = "dm"
		}
		if choiceint == 3 {
			outputint = numberint * 10
			unit = "cm"
		}
		if choiceint == 4 {
			outputint = numberint * 10
			unit = "dm"
		}
		if choiceint == 5 {
			outputint = numberint * 0.1
			unit = "m"
		}
		if choiceint == 6 {
			outputint = numberint / 60
			unit = "minutes"
		}
		if choiceint == 7 {
			outputint = numberint / 3600
			unit = "hours"
		}
		if choiceint == 8 {
			outputint = numberint * 60
			unit = "seconds"
		}
		if choiceint == 9 {
			outputint = numberint * 3600
			unit = "seconds"
		}
		if choiceint == 10 {
			outputint = numberint / 60
			unit = "hours"
		}
		if choiceint == 11 {
			outputint = numberint * 60
			unit = "minutes"
		}
		if choiceint == 12 {
			outputint = numberint / 1000
			unit = "kg"
		}
		if choiceint == 13 {
			outputint = numberint * 1000
			unit = "g"
		}
		if choiceint == 14 {
			outputint = numberint / 1000
			unit = "l"
		}
		if choiceint == 15 {
			outputint = numberint * 1000
			unit = "ml"
		}

		Clear()
		fmt.Printf("It's %f %s\n", outputint, unit)
	}
}
