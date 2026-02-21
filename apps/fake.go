package apps

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Main15() {
	Clear()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	welcometext := "----Welcome to FakeLogGen----"
	minuses := (len(welcometext)/2 - len(" Bye! ")/2)
	max := 18
	min := 2
	filenames := []string{"hack.log", "exploit.py.log", "exploit2.log", "exploit.log", "hack2.log"}
	errors := []string{
		"[FAILED] Can't execute exploit",
		"[FAILED] Deauth can't be executed: No WiFi module",
		"[FAILED] RamOverload failed",
		"[SUCESS] Sucessfully Deauthed",
		"[SUCESS] Exploit Executed",
		"[SUCESS] OverLoaded RAM",
	}
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				cancel()
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(2 * time.Second)
			Clear()
			return
		default:
		}

		filename := filenames[rand.Intn(len(filenames))]
		fmt.Println("-----" + filename + "-----")

		errors2 := ""
		for i := 0; i < rand.Intn(max-min+1)+min; i++ {
			errors2 += errors[rand.Intn(len(errors))] + "\n"
		}

		typeeffect(ctx, errors2, filename, welcometext)
		time.Sleep(time.Second * 2)
	}
}
