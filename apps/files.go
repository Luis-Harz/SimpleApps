package apps

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

var dir string
var recursive bool

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
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
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

func nano(file string) {
	Clear()
	for {
		lines := readfile(file)
		fmt.Printf("[%s; %d Lines]%s ", file, len(lines), ConfigData.Prompt)
		input := ReadInput()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		argument := ""
		if len(parts) > 1 {
			argument = parts[1]
		}
		if command == "exit" {
			break
		}
		if command == "file" && argument != "" {
			if argument == "show" {
				printfile(lines)
			} else if argument == "edit" {
				fmt.Println("Line: ")
				fmt.Printf("%s ", ConfigData.Prompt)
				input := ReadInput()
				line, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println(err)
					time.Sleep(time.Second * 2)
					continue
				}
				fmt.Println("Content: ")
				fmt.Printf("%s ", ConfigData.Prompt)
				input = ReadInput()
				for len(lines) <= line {
					lines = append(lines, "")
				}
				lines[line] = input
				saveFile(lines, file)
			} else {
				fmt.Println("Invalid argument!")
			}
		} else if command == "clear" {
			Clear()
		} else {
			fmt.Println("Command not found!")
		}
	}
}

func Main21() {
	for {
		recursive = false
		_, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Failed to get Term Size")
			time.Sleep(time.Second * 2)
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
			fmt.Println("Invalid Command")
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
						fmt.Println(err)
						time.Sleep(time.Second * 2)
						continue
					}
				}
			} else if command == "cf" {
				file, err := os.Create(argument)
				if err != nil {
					fmt.Println("Error creating file", err)
					time.Sleep(2 * time.Second)
					continue
				}
				file.Close()
			} else if command == "ef" {
				if argument == "" {
					fmt.Println("No file specified!")
					continue
				}
				nano(argument)
			} else if command == "mkd" {
				os.Mkdir(argument, 0755)
			} else if command == "cd" {
				if argument == "" {
					fmt.Println("No directory specified!")
					continue
				}
				err := os.Chdir(argument)
				if err != nil {
					fmt.Println("Error changing directory:", err)
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
			}
		}
	}
}
