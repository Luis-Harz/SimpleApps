package apps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Main5() {
	Clear()
	welcometext := "----Welcome to CoinFlip----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	frame1 := `
###########
###########
##### #####
##### #####
##### #####
###########
###########
	`
	frame2 := `
###########
###########
#### # ####
#### # ####
#### # ####
###########
###########
	`
	frame := false
	var choosen string
	var flipped string
	var coinflipped int
	var coinflipped2 int
	for {
		coinflipped = 0
		fmt.Println("[Type C to start/continue or E to exit]")
		fmt.Printf("%s ", ConfigData.Prompt)
		fmt.Scanln(&choosen)
		if choosen == "E" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		} else if choosen == "C" {
			Clear()
			flips := Random(5, 20)
			for i := 0; i < flips; i++ {
				frame = !frame
				if frame == true {
					fmt.Println(frame1)
				} else {
					fmt.Println(frame2)
				}
				time.Sleep(400 * time.Millisecond)
				coinflipped++
				coinflipped2++
				Clear()
			}
			if frame == true {
				flipped = "I"
			} else {
				flipped = "II"
			}
			fmt.Println("You flipped....." + flipped)
			fmt.Println("The Coin flipped " + strconv.Itoa(coinflipped))
			fmt.Println("The Coin flipped " + strconv.Itoa(coinflipped2) + " in total")
		} else {
			fmt.Println("Please give a valid input!")
			continue
		}
	}
}
