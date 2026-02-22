package apps

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

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
	printfile(readfile(file))
}
func showfilecontentnanopre(content []string) []string {
	Clear()
	aftercontent := []string{}
	for i := 0; i < len(content); i++ {
		fmt.Println(content[i])
		aftercontent = append(aftercontent, content[i])
	}
	drawnano(y, x)
	return aftercontent
}

var x int
var y int
var filebefore []string
var ttyClosed bool = false

func savefilenano(lines []string, file string, tty tty.TTY) {
	if !ttyClosed {
		tty.Close()
		ttyClosed = true
	}
	fileafter := showfilecontentnanopre(lines)
	if !reflect.DeepEqual(filebefore, fileafter) {
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
		}
	}
}

func nano(file string) {
	showfilecontentnano(file)
	lines := readfile(file)
	filebefore = readfile(file)
	x = 1
	y = 1
	tty, err := tty.Open()
	if err != nil {
		showerror(err.Error())
	}
	defer func() {
		if !ttyClosed {
			tty.Close()
			ttyClosed = true
		}
	}()
	for {
		drawnano(y, x)
		r, err := tty.ReadRune()
		if err != nil {

			showerror(err.Error())
		}

		switch r {
		case 0x08, 0x7f:
			if x > 1 {
				x--
				lines[y-1] = lines[y-1][:x-1] + lines[y-1][x:]
				//saveFile(lines, file)
				//showfilecontentnano(file)
				showfilecontentnanopre(lines)
			}
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
					if y < len(lines) {
						y++
					}
				case 'C':
					if x < (len(lines[(y-1)]) + 1) {
						x++
					} else if y < len(lines) {
						x = 1
						y++
					}
				case 'D':
					if x > 1 {
						x--
					} else if y > 1 {
						y--
						x = len(lines[y-1]) + 1
					}
				}
			}
		case '\r', '\n':
			newLine := ""
			if x <= len(lines[y-1]) {
				newLine = lines[y-1][x-1:]
				lines[y-1] = lines[y-1][:x-1]
			}
			lines = append(lines[:y], append([]string{newLine}, lines[y:]...)...)
			y++
			x = 1
			showfilecontentnanopre(lines)
		case 'q':
			savefilenano(lines, file, *tty)
			return
		default:
			if x > len(lines[y-1]) {
				lines[y-1] += string(r)
			} else {
				lines[y-1] = lines[y-1][:x-1] + string(r) + lines[y-1][x-1:]
			}
			x++
			showfilecontentnanopre(lines)
		}
	}
}

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
