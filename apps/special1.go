package apps

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"

	"golang.org/x/term"
)

func Main10() {

	for {
		fmt.Print("\033[H\033[2J") // clear

		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			return
		}

		text := "800+ Lines of Code"
		textLen := utf8.RuneCountInString(text)

		row := height / 2
		col := (width - textLen) / 2
		if col < 0 {
			col = 0
		}

		for r := 0; r < height; r++ {
			line := make([]rune, width)
			for i := 0; i < width; i++ {
				line[i] = rune(randChar())
			}
			if r == row {
				for i, ch := range text {
					if col+i < width {
						line[col+i] = ch
					}
				}
			}

			fmt.Println(string(line))
		}

		time.Sleep(400 * time.Millisecond)
	}
}
