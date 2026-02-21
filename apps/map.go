package apps

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func Main13() {
	Clear()
	welcometext := "----Welcome to Map Gen----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	for {
		var input string
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Failed to get Term Size")
			time.Sleep(time.Second * 2)
			Clear()
		}
		fmt.Print(" [0] Generate\n [1] Read\n [2] Exit\n")
		fmt.Print(ConfigData.Prompt + " ")
		fmt.Scan(&input)
		if input == "2" || input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			Clear()
			break
		}
		inputint, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if inputint == 0 {
			var mapped string
			width, height, _ = term.GetSize(int(os.Stdout.Fd()))
			mapped = fmt.Sprintf("%d %d\n", width, height)
			for i := 0; i < width*height; i++ {
				if rand.Intn(2) == 1 {
					mapped += "0"
				} else {
					mapped += "1"
				}
			}
			fmt.Print(decode(mapped))
			time.Sleep(time.Second * 5)
			Clear()
			var filename string
			fmt.Println("Filename(without ending): ")
			fmt.Print(ConfigData.Prompt + " ")
			fmt.Scan(&filename)
			filename = filename + ".map"
			err = os.WriteFile(filename, []byte(mapped), 0644)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Print("Saved!")
			time.Sleep(3 * time.Second)
			Clear()
		} else if inputint == 1 {
			var filename string
			fmt.Println("Filename(without ending): ")
			fmt.Print(ConfigData.Prompt + " ")
			fmt.Scan(&filename)
			filename = filename + ".map"
			data, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			content := string(data)
			lines := strings.SplitN(content, "\n", 2)
			if len(lines) < 2 {
				fmt.Println("Invalid map file")
				time.Sleep(2 * time.Second)
				return
			}
			var mapWidth, mapHeight int
			fmt.Sscanf(lines[0], "%d %d", &mapWidth, &mapHeight)
			mainContent := lines[1]
			termWidth, termHeight, _ := term.GetSize(int(os.Stdout.Fd()))
			if termWidth < mapWidth || termHeight < mapHeight {
				fmt.Println("Terminal too small, can't load map!")
				time.Sleep(2 * time.Second)
				return
			}
			if len(mainContent) < mapWidth*mapHeight {
				fmt.Println("Terminal too big. Problems may appear")
				time.Sleep(time.Second * 2)
				Clear()
				mainContent += strings.Repeat("-", mapWidth*mapHeight-len(mainContent))
			} else if len(mainContent) > mapWidth*mapHeight {
				mainContent = mainContent[:mapWidth*mapHeight]
			}
			mapped := insertbreakEverywidth(mainContent, mapWidth)
			decoded := decode(mapped)
			fmt.Print(decoded)
			time.Sleep(time.Second * 3)
		}
	}
}
