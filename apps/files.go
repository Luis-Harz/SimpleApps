package apps

import (
	"bufio"
	"fmt"
	"os"

	//"reflect"

	//"strconv"
	"strings"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

var dir string
var recursive bool

func showHelp(lines []string) {
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return
	}
	start := height - len(lines) - 1
	for i, line := range lines {
		fmt.Printf("\033[%d;1H\033[K%s", start+i, line)
	}
	fmt.Printf("\033[%d;1H", height)
}

func showerror(error string) {
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Failed to get Term Size")
		time.Sleep(time.Second * 2)
	}
	fmt.Printf("\033[%d;1H%s\033[%d;1H", (height - 2), error, height)
}

func listfiles() {
	cwd, _ := os.Getwd()
	entries, _ := os.ReadDir(cwd)
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			name += "/"
		}
		fmt.Println(name)
	}
}

func readfile(infile string) []string {
	file, err := os.Open(infile)
	if err != nil {
		showerror(err.Error())
		time.Sleep(time.Second)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		showerror(err.Error())
		time.Sleep(time.Second)
	}
	return lines
}

func saveFile(lines []string, filename string) error {
	data := strings.Join(lines, "\n")
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func printfile(parts []string) {
	for i := 0; i < len(parts); i++ {
		fmt.Println(parts[i])
	}
}

func drawnano(y int, x int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func showfilecontentnano(file string) {
	Clear()
	//printfile(readfile(file))
	lines := readfile(file)
	for i := minline; i < maxline; i++ {
		fmt.Println(lines[i])
	}
}

func showfilecontentnanopre(content []string) {
	Clear()
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))
	windowSize := height - 1
	for i := 0; i < windowSize; i++ {
		lineIndex := minline + i
		if lineIndex < len(content) {
			fmt.Println(content[lineIndex])
		}
	}
	drawnano(cursorY+1, x)
}

var x int = 1
var y int = 1
var filebefore []string = []string{}
var ttyClosed bool = false

func savefilenano(lines []string, file string, tty tty.TTY) {
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

func showfilewindow(lines []string) {
	Clear()
	_, height, _ := term.GetSize(int(os.Stdout.Fd()))
	windowSize := height - 1
	for i := 0; i < windowSize; i++ {
		idx := minline + i
		if idx < len(lines) {
			fmt.Println(lines[idx])
		} else {
			fmt.Println()
		}
	}
	drawnano(cursorY+1, x)
}

func nano(file string) {
	lines := readfile(file)
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
		showfilewindow(lines)
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
			savefilenano(lines, file, *tty)
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

var minline int = 0
var maxline int
var cursorY int = 0

func Main21() {
	Clear()
	Greet("SimpleFiles")
	fmt.Println("Press Enter to start")
	ReadInput()
	for {
		recursive = false
		_, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			showerror("Can't get Terminal info!")
			time.Sleep(time.Second)
			continue
		}
		dir, _ = os.Getwd()
		Clear()
		fmt.Printf("---%s---\n", dir)
		listfiles()
		fmt.Printf("\033[%d;1H", height)
		fmt.Printf("%s ", ConfigData.Prompt)
		input := ReadInput()
		parts := strings.Fields(input)
		if len(parts) > 3 {
			showerror("Command not found!")
			time.Sleep(time.Second)
			continue
		} else {
			var argument string
			var command string
			if len(parts) > 0 {
				command = parts[0]
			}
			if len(parts) > 1 {
				argument = parts[1]
			}
			if len(parts) == 3 && parts[1] == "-r" {
				recursive = true
				argument = parts[2]
				command = parts[0]
			}
			if command == "rf" {
				if recursive == true {
					os.RemoveAll(argument)
				} else {
					err := os.Remove(argument)
					if err != nil {
						showerror(err.Error())
						time.Sleep(time.Second)
						continue
					}
				}
			} else if command == "cf" {
				file, err := os.Create(argument)
				if err != nil {
					showerror("Error: " + err.Error())
					time.Sleep(time.Second)
					time.Sleep(2 * time.Second)
					continue
				}
				file.Close()
			} else if command == "ef" {
				if argument == "" {
					showerror("No file specified!")
					time.Sleep(time.Second * 1)
					continue
				}
				nano(argument)
			} else if command == "mkd" {
				os.Mkdir(argument, 0755)
			} else if command == "cd" {
				if argument == "" {
					showerror("No Directory specified")
					time.Sleep(time.Second)
					continue
				}
				err := os.Chdir(argument)
				if err != nil {
					showerror("Error: " + err.Error())
					time.Sleep(time.Second)
					time.Sleep(2 * time.Second)
					continue
				}
			} else if command == "rd" {
				if recursive == true {
					os.RemoveAll(argument)
				} else {
					err := os.Remove(argument)
					if err != nil {
						fmt.Println(err)
						time.Sleep(time.Second * 2)
						continue
					}
				}
			} else if command == "exit" {
				Bye()
				return
			} else if command == "help" {

				commands := []string{"rf", "cf", "ef", "mkd", "cd", "rd"}
				meanings := []string{
					"Remove file",
					"Create file",
					"Edit file",
					"Make dir",
					"Change dir",
					"Remove dir",
				}

				var lines []string
				lines = append(lines, "--- Commands ---")

				for i := 0; i < len(commands); i++ {
					lines = append(lines,
						fmt.Sprintf(" [%d] %s; %s", i, commands[i], meanings[i]),
					)
				}

				showHelp(lines)
				time.Sleep(time.Second * 3)
			} else {
				showerror("Command not found! type 'help' to see all commands!")
				time.Sleep(time.Second * 2)
			}
		}
	}
}
