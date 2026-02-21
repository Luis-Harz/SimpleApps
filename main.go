package main

import (
	"SimpleApps/apps"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func menu() {
	programs := []func(){apps.Main1, apps.Main2, apps.Main3, apps.Main4, apps.Main5, apps.Main6, apps.Main7, apps.Main8, apps.Main9, apps.Main10, apps.Main11, apps.Main12, apps.Main13, apps.Main14, apps.Main15, apps.Main16, apps.Main17, apps.Main18, apps.Main19, apps.Main20, apps.Main21, apps.Update}
	names := []string{"NumberChecker", "GradeChecker", "UnitConverter", "Number2Bar", "CoinFlip", "Countdown", "Timer", "Clock", "Magic 8-Ball", "800+ Lines Special", "Calculator", "ToDo List", "Map Gen", "Matrix", "FakeLogGen", "SysMonitor", "ClockV2", "ASCII Animations", "SimpleChat", "PasswordGen", "SimpleFiles", "Update"}
	rand.Seed(time.Now().UnixNano())
	apps.Configure()
	data, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println("Error reading local version:", err)
		time.Sleep(2 * time.Second)
		return
	}
	localVersion := strings.TrimSpace(string(data))
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "help" {
			fmt.Println("--Tools--")
			for i := 0; i < len(names); i++ {
				fmt.Printf(" [%d] %s\n", i, names[i])
			}
		} else {
			argint, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Print("Tool not found. Type 'help' to show all tools")
			} else {
				if argint > -1 && argint < len(programs)+1 {
					programs[argint]()
				}
			}
		}
	} else {
		for {
			apps.Clear()
			welcometext := "----Welcome to SimpleApps----"
			minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
			fmt.Println(welcometext)
			fmt.Println(strings.Repeat("-", 6) + "You are on V" + localVersion + strings.Repeat("-", 6))
			fmt.Println("What do you want to run?")
			for i, name := range names {
				fmt.Printf("[%d] %s\n", i, name)
			}
			fmt.Printf("[%d] Exit\n", len(programs))
			fmt.Printf("%s ", apps.ConfigData.Prompt)
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
				time.Sleep(3 * time.Second)
				apps.Clear()
				break
			}

			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid choice, try again!")
				continue
			}

			if choice == len(programs) {
				fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
				time.Sleep(3 * time.Second)
				apps.Clear()
				break
			} else if choice >= 0 && choice < len(programs) {
				programs[choice]()
			} else {
				fmt.Println("Invalid choice, try again!")
				apps.Clear()
			}
		}
	}
}

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting Terminal size:", err)
		return
	}
	fmt.Printf("Terminal: %d columns, %d lines\n", width, height)

	minWidth := 70
	minHeight := 23

	if width < minWidth || height < minHeight {
		fmt.Printf("[Warning] Terminal is too small, min size: %dx%d\n", minWidth, minHeight)
	} else {
		menu()
	}
}
