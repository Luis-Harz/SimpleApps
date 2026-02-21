package apps

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

func Main14() {
	green := "\033[32m"
	reset := "\033[0m"
	Clear()
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting Terminal size:", err)
		return
	}
	for {
		Softclear()
		string := ""
		for i := 0; i < width*height; i++ {
			if rand.Intn(2) == 0 {
				string += "0"
			} else {
				string += "1"
			}
		}
		fmt.Print(green + string + reset)
		time.Sleep(100 * time.Millisecond)
	}
}
