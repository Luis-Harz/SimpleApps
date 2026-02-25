package apps

import (
	"fmt"
	"os"
	"strings"

	//"strconv"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

//editor

func savefilepic(lines []string, tty tty.TTY) {
	if !ttyClosed {
		tty.Close()
		ttyClosed = true
	}
	//fileafter := showfilecontentnanopre(lines)
	//if !reflect.DeepEqual(filebefore, fileafter) {
	for {
		Clear()
		fmt.Println("Do you want to save?")
		fmt.Printf("%s ", ConfigData.Prompt)
		input := ReadInput()
		if input == "y" {
			file := Input("File Name?")
			saveFile(lines, file)
			break
		} else if input == "n" {
			return
		} else {
			continue
		}
		//}
	}
}

func showfilewindowpic(lines []string) {
	Clear()
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))
	windowSize := height - 1

	for i := 0; i < windowSize; i++ {
		idx := minline + i
		if idx < len(lines) {
			line := lines[idx]

			coloredLine := ""
			for _, ch := range line {
				if ch >= '0' && ch <= '9' {
					number := int(ch - '0')
					if number >= 0 && number < len(variables)-1 {
						coloredLine += variables[number] + string(ch) + reset
					} else {
						coloredLine += string(ch)
					}
				} else {
					coloredLine += string(ch)
				}
			}

			fmt.Println(coloredLine)
		} else {
			fmt.Println()
		}
	}

	drawnano(cursorY+1, x)
}

func piceditor() {
	x = 1
	y = 1
	filebefore = []string{}
	ttyClosed = false
	lines := []string{}
	if len(lines) == 0 {
		lines = append(lines, "")
	}
	filebefore = append([]string{}, lines...)
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))
	windowSize := height - 1
	minline = 0
	cursorY = 0
	x = 1

	tty, err := tty.Open()
	if err != nil {
		showerror(err.Error())
		return
	}

	for {
		showfilewindowpic(lines)
		drawnano(cursorY+1, x)
		r, _ := tty.ReadRune()
		curIndex := minline + cursorY
		if curIndex >= len(lines) {
			lines = append(lines, "")
		}
		line := lines[curIndex]

		switch r {
		case 0x08, 0x7f:
			if x > 1 && x-1 <= len(line) {
				lines[curIndex] = line[:x-2] + line[x-1:]
				x--
			} else if x == 1 && curIndex > 0 {
				prev := lines[curIndex-1]
				lines[curIndex-1] = prev + line
				lines = append(lines[:curIndex], lines[curIndex+1:]...)
				if cursorY > 0 {
					cursorY--
				} else if minline > 0 {
					minline--
				}
				x = len(prev) + 1
			}
		case 27:
			seq := []rune{0, 0}
			seq[0], _ = tty.ReadRune()
			seq[1], _ = tty.ReadRune()
			if seq[0] == '[' {
				switch seq[1] {
				case 'A':
					if cursorY > 0 {
						cursorY--
					} else if minline > 0 {
						minline--
					}
					if x > len(lines[minline+cursorY])+1 {
						x = len(lines[minline+cursorY]) + 1
					}
				case 'B':
					if cursorY < windowSize-1 && minline+cursorY < len(lines)-1 {
						cursorY++
					} else if minline+windowSize < len(lines) {
						minline++
					}
					if x > len(lines[minline+cursorY])+1 {
						x = len(lines[minline+cursorY]) + 1
					}
				case 'C':
					if x <= len(lines[curIndex]) {
						x++
					} else if curIndex < len(lines)-1 {
						x = 1
						if cursorY < windowSize-1 {
							cursorY++
						} else {
							minline++
						}
					}
				case 'D':
					if x > 1 {
						x--
					} else if curIndex > 0 {
						cursorY--
						if cursorY < 0 {
							minline--
							cursorY = 0
						}
						x = len(lines[minline+cursorY]) + 1
					}
				}
			}
		case '\r', '\n':
			newLine := ""
			if x <= len(line) {
				newLine = line[x-1:]
				lines[curIndex] = line[:x-1]
			}
			lines = append(lines[:curIndex+1], append([]string{newLine}, lines[curIndex+1:]...)...)
			if cursorY < windowSize-1 {
				cursorY++
			} else {
				minline++
			}
			x = 1
		case 'q':
			if ttyClosed != true {
				tty.Close()
			}
			Clear()
			savefilepic(lines, *tty)
			return
		default:
			if x > len(line) {
				line += string(r)
			} else {
				line = line[:x-1] + string(r) + line[x-1:]
			}
			lines[curIndex] = line
			x++
		}
	}
}

//END editor

// colors
var red string = "\033[31m"
var blue string = "\033[34m"
var green string = "\033[32m"
var white string = "\033[37m"
var black string = "\033[30m"
var reset string = "\033[0m"
var newline string = "\n"
var block string = "██"
var variables []string = []string{red, green, blue, white, black, newline}

func getpicwidth(picture []int) int {
	var width int
	for i := 0; i < len(picture); i++ {
		if picture[i] != len(variables)-1 {
			continue
		} else {
			width = i
			break
		}
	}
	return width * 2
}

func showpic(picture []int) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting Terminal size:", err)
		return
	}
	if getpicwidth(picture) > width {
		Warn("Image too big to display")
		time.Sleep(time.Second * 2)
	} else {
		for i := 0; i < len(picture); i++ {
			if picture[i] != len(variables)-1 {
				fmt.Print(variables[picture[i]] + block + reset)
			} else {
				fmt.Print(newline)
			}
		}
	}
}

func Main25() {
	Clear()
	features := []string{"Make picture\n", "Display picture\n"}
	fmt.Println("---Functions---")
	for i := 0; i < len(features); i++ {
		fmt.Printf(" [%d] %s", i, features[i])
	}
	input := Input("What do you want to do?")
	if input == "0" {
		piceditor()
	} else if input == "1" {
		filename := Input("File Name:")
		filecontent, err := os.ReadFile(filename)
		content := strings.ReplaceAll(string(filecontent), "\r", "")
		filecontent = []byte(content)
		if err != nil {
			showerror(err.Error())
		}
		for {
			Clear()
			picture := []int{}
			for i := 0; i < len(filecontent); i++ {
				ch := filecontent[i]

				if ch == '\n' {
					picture = append(picture, len(variables)-1)
					continue
				}

				if ch == '\r' {
					continue
				}

				if ch >= '0' && ch <= '9' {
					picture = append(picture, int(ch-'0'))
				}
			}
			showpic(picture)
			input := Input("\n[Type 'exit' to stop viewing]")
			if input == "exit" {
				return
			} else {
				continue
			}
		}
	}
}
