package apps

import (
	"bufio"
	"fmt"
	"os"

	//"reflect"
	"strconv"
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
	time.Sleep(time.Second * 2)
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
		var parts2 []string
		flags := make(map[string]bool)
		for _, part := range parts {
			if strings.HasPrefix(part, "-") {
				flags[part] = true
			} else {
				parts2 = append(parts2, part)
			}
		}
		command := parts2[0]
		args := parts2[1:]
		recursive := flags["-r"]
		if len(parts) > 3 {
			showerror("Command not found!")
			time.Sleep(time.Second)
			continue
		} else {
			if command == "rm" {
				if len(args) > 0 {
					if recursive == true {
						err := os.RemoveAll(args[0])
						if err != nil {
							showerror(err.Error())
						}
					} else {
						err := os.Remove(args[0])
						if err != nil {
							showerror(err.Error())
						}
					}
				} else {
					showerror("No argument given!")
				}
			} else if command == "edit" {
				if len(args) > 0 {
					nano(args[0])
				} else {
					showerror("No file given!")
				}
			} else if command == "cf" {
				if len(args) > 0 {
					file, err := os.Create(args[0])
					if err != nil {
						showerror(err.Error())
					}
					file.Close()
				} else {
					showerror("No file given!")
				}
			} else if command == "mkd" {
				if len(args) > 0 {
					if recursive == true {
						if len(args) > 2 {
							filemode, err := strconv.Atoi(args[1])
							if err != nil {
								showerror(err.Error())
							}
							os.MkdirAll(args[0], os.FileMode(filemode))
						} else {
							os.MkdirAll(args[0], 0755)
						}
					} else {
						if len(args) > 2 {
							filemode, err := strconv.Atoi(args[1])
							if err != nil {
								showerror(err.Error())
							}
							os.Mkdir(args[0], os.FileMode(filemode))
						} else {
							os.Mkdir(args[0], 0755)
						}
					}
				} else {
					showerror("No dir given")
				}
			} else if command == "cd" {
				if len(args) > 0 && len(args) < 2 {
					os.Chdir(args[0])
				}
			} else if command == "help" {
				commands := []string{"rm", "edit", "cf", "mkd", "cd"}
				functions := []string{"Remove file/dir", "edit file", "create file", "make dir", "change dir"}
				functions2 := make([]string, len(commands))
				for i := 0; i < len(commands); i++ {
					line := fmt.Sprintf("[%d] %s; %s", i, commands[i], functions[i])
					functions2 = append(functions2, line)
				}
				showHelp(functions2)
				time.Sleep(time.Second * 2)
			}
		}
	}
}
