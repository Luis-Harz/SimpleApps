package apps

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

func draw(y int, x int, symbol string) {
	fmt.Print("\033[2J")
	fmt.Printf("\033[%d;%dH", y, x)
	fmt.Print(symbol)
}

func Main22() {
	var editedinput byte
	fmt.Println("What Symbol do you want to be?")
	fmt.Printf("%s ", ConfigData.Prompt)
	input := ReadInput()
	fmt.Println("Press Arrowkeys to move")
	if len(input) > 1 {
		editedinput = input[0]
	}
	game(string(editedinput))
}

func game(symbol string) {
	x := 1
	y := 2
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	fmt.Print("\033[?25l")
	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch r {
		case 27:
			seq := make([]rune, 2)
			for i := 0; i < 2; i++ {
				seq[i], _ = tty.ReadRune()
			}
			if seq[0] == '[' {
				switch seq[1] {
				case 'A':
					if y > 1 {
						y--
					}
				case 'B':
					if y < height {
						y++
					}
				case 'C':
					if x < width {
						x++
					}
				case 'D':
					if x > 1 {
						x--
					}
				}
			}
		case 'q':
			fmt.Print("\033[?25h")
			return
		}
		draw(y, x, symbol)
	}
}
