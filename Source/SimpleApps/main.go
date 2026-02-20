package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/websocket"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/term"
)

var filetodownload string = ""
var serverurl string = "simplemirror.bolucraft.uk/ForServer"
var config Config

// ASCII Digits START
var digits = map[rune]string{
	'0': `
 ### 
#   #
#   #
#   #
 ### 
`,
	'1': `
  #  
 ##  
  #  
  #  
 ### 
`,
	'2': `
 ### 
    #
 ### 
#    
 ### 
`,
	'3': `
 ### 
    #
 ### 
    #
 ### 
`,
	'4': `
#   #
#   #
 ### 
    #
    #
`,
	'5': `
 ### 
#    
 ### 
    #
 ### 
`,
	'6': `
 ### 
#    
 ### 
#   #
 ### 
`,
	'7': `
 ### 
    #
   # 
  #  
  #  
`,
	'8': `
 ### 
#   #
 ### 
#   #
 ### 
`,
	'9': `
 ### 
#   #
 ### 
    #
 ### 
`,
	':': `
     
  #  
     
  #  
     
`,
}

//END

// ASCII Digits V2 START
var digits2 = map[rune]string{
	'0': `
█████
█   █
█   █
█   █
█████
`,
	'1': `
  █  
  █  
  █  
  █  
  █  
`,
	'2': `
█████
    █
█████
█    
█████
`,
	'3': `
█████
    █
█████
    █
█████
`,
	'4': `
█   █
█   █
█████
    █
    █
`,
	'5': `
█████
█    
█████
    █
█████
`,
	'6': `
█████
█    
█████
█   █
█████
`,
	'7': `
█████
    █
    █
    █
    █
`,
	'8': `
█████
█   █
█████
█   █
█████
`,
	'9': `
█████
█   █
█████
    █
█████
`,
	':': `
     
  █  
     
  █  
     
`,
}

//END

// Random Number
func random(minimum int, maximum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(maximum-minimum+1) + minimum
	return n
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func softclear() {
	fmt.Print("\033[H")
}

func downloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if filename == "SimpleApps" || filename == "SimpleApps.exe" {
		os.Chmod(filename, 0755)
	}
	return err
}

func update() {
	clear()
	urlVersion := "http://" + serverurl + "/version.txt"
	data, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println("Error reading local version:", err)
		time.Sleep(time.Second * 1)
		return
	}
	localVersion := strings.TrimSpace(string(data))
	resp, err := http.Get(urlVersion)
	if err != nil {
		fmt.Println("Error fetching remote version:", err)
		time.Sleep(time.Second * 1)
		return
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading remote version:", err)
		time.Sleep(time.Second * 1)
		return
	}
	remoteVersion := strings.TrimSpace(string(body))
	fmt.Println("Local Version:", localVersion, "=> Remote Version:", remoteVersion)
	if localVersion == remoteVersion {
		fmt.Println("Already up-to-date!")
		time.Sleep(time.Second * 1)
		return
	}
	fmt.Println("Do you want to upgrade?\n [0] Yes\n [1] No")
	fmt.Print(config.Prompt + " ")
	var input string
	fmt.Scanln(&input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice != 0 {
		fmt.Println("Update canceled.")
		time.Sleep(time.Second * 1)
		return
	}
	var fileName string
	if runtime.GOOS == "windows" {
		fileName = "SimpleApps.exe"
	} else {
		fileName = "SimpleApps"
	}

	tmpFile := fileName + ".tmp"
	fmt.Println("Downloading new version...")
	err = downloadFile("http://"+serverurl+"/"+fileName, tmpFile)
	if err != nil {
		fmt.Println("Download failed:", err)
		time.Sleep(time.Second * 1)
		return
	}
	if runtime.GOOS != "windows" {
		os.Chmod(tmpFile, 0755)
		os.Rename(tmpFile, fileName)
		fmt.Println("Updated successfully!")
	} else {
		fmt.Println("Windows: Please restart the application to complete the update.")
	}
	err = downloadFile("http://"+serverurl+"/version.txt", "version.txt")
	if err != nil {
		fmt.Println("Failed to update version.txt:", err)
		time.Sleep(time.Second * 1)
	} else {
		fmt.Println("Version file updated!")
	}
	time.Sleep(time.Second * 1)
}

func main1() {
	clear()
	welcometext := "----Welcome to NumberChecker----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	for {
		var input string
		fmt.Println("Type a Number(or 'exit' to exit): ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please Type a valid Number!")
			continue
		}
		clear()
		if number%2 == 0 {
			fmt.Printf("Number %d is even\n", number)
		} else {
			fmt.Printf("Number %d is odd\n", number)
		}
	}
}

func main2() {
	clear()
	welcometext := "----Welcome to GradeChecker----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	for {
		var maxscorestring string
		var scoreStr string
		fmt.Println("What's the Maximal Score(type 'exit' to exit)? ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&maxscorestring)
		if maxscorestring == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		maxscorestring = strings.TrimSpace(maxscorestring)
		maxscorestring = strings.ReplaceAll(maxscorestring, ",", ".")
		fmt.Println("What's the Student's Score? ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&scoreStr)
		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			fmt.Println("Please type a valid number!")
			continue
		}
		maxscoreFloat, err := strconv.ParseFloat(maxscorestring, 64)
		if err != nil {
			fmt.Println("Invalid max score!")
			continue
		}
		percent := score / maxscoreFloat * 100
		clear()
		switch {
		case percent >= 90:
			fmt.Printf("With a Score of %.1f the Student get's an A\n", score)
		case percent >= 80:
			fmt.Printf("With a Score of %.1f the Student get's a B\n", score)
		case percent >= 70:
			fmt.Printf("With a Score of %.1f the Student get's a C\n", score)
		case percent >= 60:
			fmt.Printf("With a Score of %.1f the Student get's a D\n", score)
		default:
			fmt.Printf("With a Score of %.1f the Student get's an F\n", score)
		}
	}
}

func main3() {
	clear()
	welcometext := "----Welcome to UnitConverter----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)

	for {
		// Length
		fmt.Println(" [0]  cm => m")
		fmt.Println(" [1]  m => cm")
		fmt.Println(" [2]  cm => dm")
		fmt.Println(" [3]  dm => cm")
		fmt.Println(" [4]  m => dm")
		fmt.Println(" [5]  dm => m")
		// Time
		fmt.Println(" [6]  seconds => minutes")
		fmt.Println(" [7]  seconds => hours")
		fmt.Println(" [8]  minutes => seconds")
		fmt.Println(" [9]  hours => seconds")
		fmt.Println(" [10] minutes => hours")
		fmt.Println(" [11] hours => minutes")
		// Mass
		fmt.Println(" [12] g => kg")
		fmt.Println(" [13] kg => g")
		// Volume
		fmt.Println(" [14] ml => l")
		fmt.Println(" [15] l => ml")

		var choice string
		fmt.Println("What Conversion (or type 'exit' to exit)?")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}

		choiceint, err := strconv.Atoi(choice)
		if err != nil {
			clear()
			fmt.Println("Please type a valid Number!")
			continue
		}

		var number string
		fmt.Println("How much?")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&number)
		numberint, err := strconv.ParseFloat(number, 64)
		if err != nil {
			clear()
			fmt.Println("Please type a valid Number!")
			continue
		}

		var outputint float64
		var unit string

		if choiceint == 0 {
			outputint = numberint * 0.01
			unit = "m"
		}
		if choiceint == 1 {
			outputint = numberint * 100
			unit = "cm"
		}
		if choiceint == 2 {
			outputint = numberint * 0.1
			unit = "dm"
		}
		if choiceint == 3 {
			outputint = numberint * 10
			unit = "cm"
		}
		if choiceint == 4 {
			outputint = numberint * 10
			unit = "dm"
		}
		if choiceint == 5 {
			outputint = numberint * 0.1
			unit = "m"
		}
		if choiceint == 6 {
			outputint = numberint / 60
			unit = "minutes"
		}
		if choiceint == 7 {
			outputint = numberint / 3600
			unit = "hours"
		}
		if choiceint == 8 {
			outputint = numberint * 60
			unit = "seconds"
		}
		if choiceint == 9 {
			outputint = numberint * 3600
			unit = "seconds"
		}
		if choiceint == 10 {
			outputint = numberint / 60
			unit = "hours"
		}
		if choiceint == 11 {
			outputint = numberint * 60
			unit = "minutes"
		}
		if choiceint == 12 {
			outputint = numberint / 1000
			unit = "kg"
		}
		if choiceint == 13 {
			outputint = numberint * 1000
			unit = "g"
		}
		if choiceint == 14 {
			outputint = numberint / 1000
			unit = "l"
		}
		if choiceint == 15 {
			outputint = numberint * 1000
			unit = "ml"
		}

		clear()
		fmt.Printf("It's %f %s\n", outputint, unit)
	}
}

func main4() {
	clear()
	welcometext := "----Welcome to N2B----"
	fmt.Println(welcometext)
	fmt.Println("------Number2Bar------")
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var inputstr string
	var maximalstr string
	for {
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Error getting Terminal size:", err)
			return
		}
		widthstr := strconv.Itoa(width - 2)
		fmt.Println("Please set the maximum(Terminal width: " + widthstr + ")(or type 'exit' to exit): ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&maximalstr)
		if maximalstr == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		maximal, err := strconv.Atoi(maximalstr)
		if err != nil || maximal > width {
			clear()
			fmt.Println("Please give a valid number!")
			continue
		}
		fmt.Println("Please give a number(max." + maximalstr + "): ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&inputstr)
		if inputstr == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		inputfloat, err := strconv.ParseFloat(inputstr, 64)
		if err != nil {
			clear()
			fmt.Println("Please type a valid number!")
			continue
		}
		input := int(inputfloat)
		if input < 0 || input > maximal || input > width-2 {
			clear()
			fmt.Println("Please type a valid number!")
			continue
		}
		max := maximal
		parts := max - input
		clear()
		fmt.Println("[" + strings.Repeat("#", input) + strings.Repeat(" ", parts) + "]")
	}
}

func main5() {
	clear()
	welcometext := "----Welcome to CoinFlip----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	frame1 := `
###########
###########
##### #####
##### #####
##### #####
###########
###########
	`
	frame2 := `
###########
###########
#### # ####
#### # ####
#### # ####
###########
###########
	`
	frame := false
	var choosen string
	var flipped string
	var coinflipped int
	var coinflipped2 int
	for {
		coinflipped = 0
		fmt.Println("[Type C to start/continue or E to exit]")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&choosen)
		if choosen == "E" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		} else if choosen == "C" {
			clear()
			flips := random(5, 20)
			for i := 0; i < flips; i++ {
				frame = !frame
				if frame == true {
					fmt.Println(frame1)
				} else {
					fmt.Println(frame2)
				}
				time.Sleep(400 * time.Millisecond)
				coinflipped++
				coinflipped2++
				clear()
			}
			if frame == true {
				flipped = "I"
			} else {
				flipped = "II"
			}
			fmt.Println("You flipped....." + flipped)
			fmt.Println("The Coin flipped " + strconv.Itoa(coinflipped))
			fmt.Println("The Coin flipped " + strconv.Itoa(coinflipped2) + " in total")
		} else {
			fmt.Println("Please give a valid input!")
			continue
		}
	}
}

func main6() {
	scanner := bufio.NewScanner(os.Stdin)
	welcometext := "----Welcome to Countdown----"
	minuses := (len(welcometext)/2 - len(" Bye! ")/2)

	for {
		clear()
		fmt.Println(welcometext)

		fmt.Println("How many Seconds (or type 'exit' to exit)? ")
		fmt.Printf("%s ", config.Prompt)
		if !scanner.Scan() {
			break // EOF
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(2 * time.Second)
			clear()
			break
		}

		seconds, err := strconv.Atoi(input)
		if err != nil || seconds < 0 {
			clear()
			fmt.Println("Please give a valid input!")
			time.Sleep(2 * time.Second)
			continue
		}

		for seconds > 0 {
			clear()
			h := seconds / 3600
			m := (seconds % 3600) / 60
			s := seconds % 60
			fmt.Printf("%02d:%02d:%02d\n", h, m, s)
			seconds--
			time.Sleep(time.Second)
		}

		clear()
		fmt.Println("00:00:00")
		fmt.Println("Time's Over!")
		time.Sleep(2 * time.Second)
	}
}

func main7() {
	welcometext := "----Welcome to Timer----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var start time.Time
	var paused time.Duration
	running := false

	for {
		clear()
		fmt.Println(welcometext)
		fmt.Println("Options:")
		fmt.Println("[start, stop, reset, show, exit]")

		var cmd string
		fmt.Println("Enter command: ")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&cmd)

		switch cmd {
		case "start":
			if !running {
				start = time.Now().Add(-paused)
				running = true
			}
		case "stop":
			if running {
				paused = time.Since(start)
				running = false
			}
		case "reset":
			paused = 0
			running = false
		case "show":
			if running {
				fmt.Println("Timer:", formatDuration(time.Since(start)))
			} else {
				fmt.Println("Timer:", formatDuration(paused))
			}
			time.Sleep(2 * time.Second)
		case "exit":
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			return
		default:
			fmt.Println("Unknown command!")
			time.Sleep(2 * time.Second)
		}
	}
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// Help Function main8
func printNumber(s string) {
	lines := make([][]string, len(s))
	maxLines := 0
	for i, c := range s {
		lines[i] = strings.Split(strings.Trim(digits[c], "\n"), "\n")
		if len(lines[i]) > maxLines {
			maxLines = len(lines[i])
		}
	}
	for i := 0; i < maxLines; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				fmt.Print(lines[j][i], "  ")
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

func main8() {
	clear()
	welcometext := "----Welcome to Clock----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	time.Sleep(time.Second)
	exitChan := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				exitChan <- true
				return
			}
		}
	}()
	for {
		select {
		case <-exitChan:
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			return
		default:
			clear()
			now := time.Now()
			Time := now.Format("15:04:05")
			fmt.Println("[Press 'q' or type 'exit' to exit]")
			printNumber(Time)
			fmt.Printf("%s ", config.Prompt)
			time.Sleep(time.Second)
		}
	}
}

func main9() {
	clear()
	welcometext := "----Welcome to Magic 8-Ball----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	for {
		var input string
		fmt.Println("[Type 'exit' to exit]")
		fmt.Println("Ask your question:")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		rn := rand.Intn(2)
		clear()
		if rn == 0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}

func randChar() byte {
	if rand.Intn(2) == 0 {
		return '#'
	}
	return ' '
}

func main10() {

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

/* ---------- Parser ---------- */

func (p *Parser) expr() float64 {
	result := p.term()

	for p.pos < len(p.tokens) {
		op := p.tokens[p.pos]
		if op != "+" && op != "-" {
			break
		}
		p.pos++
		if op == "+" {
			result += p.term()
		} else {
			result -= p.term()
		}
	}
	return result
}

func (p *Parser) term() float64 {
	result := p.factor()

	for p.pos < len(p.tokens) {
		op := p.tokens[p.pos]
		if op != "*" && op != "/" {
			break
		}
		p.pos++
		if op == "*" {
			result *= p.factor()
		} else {
			result /= p.factor()
		}
	}
	return result
}

type Parser struct {
	tokens []string
	pos    int
}

func (p *Parser) factor() float64 {
	token := p.tokens[p.pos]
	p.pos++

	if token == "(" {
		result := p.expr()
		p.pos++ // )
		return result
	}

	value, _ := strconv.ParseFloat(token, 64)
	return value
}

/* ---------- Tokenizer ---------- */

func tokenize(input string) []string {
	input = strings.ReplaceAll(input, " ", "")
	var tokens []string
	number := ""

	for _, c := range input {
		if (c >= '0' && c <= '9') || c == '.' {
			number += string(c)
		} else {
			if number != "" {
				tokens = append(tokens, number)
				number = ""
			}
			tokens = append(tokens, string(c))
		}
	}
	if number != "" {
		tokens = append(tokens, number)
	}
	return tokens
}

func main11() {
	clear()
	welcometext := "----Welcome to Calculator----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	var input string
	for {
		fmt.Println("[Type 'exit' to exit]")
		fmt.Printf("%s ", config.Prompt)
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		tokens := tokenize(input)
		parser := Parser{tokens: tokens}
		result := parser.expr()
		fmt.Println("=", result)
		time.Sleep(time.Second * 2)
		clear()
	}
}

// START Main12
type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// ====================Help Functions====================

func addtodo(todo string) {
	todos := loadTodos()
	newTodo := Todo{Task: todo, Completed: false}
	todos = append(todos, newTodo)
	saveTodos(todos)
	fmt.Println("Todo added!")
}

func deletetodo(todo string) {
	todos := loadTodos()
	newTodos := []Todo{}
	for _, t := range todos {
		if t.Task != todo {
			newTodos = append(newTodos, t)
		}
	}
	saveTodos(newTodos)
	fmt.Println("Todo deleted!")
}
func markCompleted(todo string) {
	todos := loadTodos()
	found := false
	for i, t := range todos {
		if t.Task == todo {
			todos[i].Completed = true
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Todo not found:", todo)
		return
	}
	saveTodos(todos)
	fmt.Println("Todo Checked!")
}

func listTodos() {
	todos := loadTodos()
	if len(todos) == 0 {
		fmt.Println("No Todos!")
		return
	}
	for _, t := range todos {
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %s\n", status, t.Task)
	}
}

// ====================JSON Load/Save====================

func loadTodos() []Todo {
	var todos []Todo
	file, err := os.Open("todos.json")
	if err != nil {
		return []Todo{}
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		return []Todo{}
	}
	return todos
}

func saveTodos(todos []Todo) {
	outFile, err := os.Create("todos.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(todos); err != nil {
		panic(err)
	}
}

// ====================Main12====================

func main12() {
	clear()
	welcometext := "----Welcome to ToDo List----"

	for {
		fmt.Println(welcometext)
		minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
		scanner := bufio.NewScanner(os.Stdin)
		listTodos()
		fmt.Printf("%s ", config.Prompt)
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()

		if strings.HasPrefix(command, "add ") {
			input := command[4:]
			addtodo(input)
		} else if strings.HasPrefix(command, "check ") {
			input := command[6:]
			markCompleted(input)
		} else if strings.HasPrefix(command, "delete ") {
			input := command[7:]
			deletetodo(input)
		} else if strings.HasPrefix(command, "help") {
			fmt.Println("Commands:\n add name\n check name\n delete name\n exit\n help")
		} else if strings.HasPrefix(command, "exit") {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		} else {
			fmt.Println("Unknown Command. Type 'help'")
		}
		time.Sleep(time.Second * 2)
		clear()
	}
}

//END Main12

// Help Function main13
func insertbreakEverywidth(s string, width int) string {
	r := []rune(s)
	var b strings.Builder

	for i, c := range r {
		b.WriteRune(c)
		if (i+1)%width == 0 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}

func decode(s string) string {
	m := map[rune]rune{
		'0': ' ',
		'1': '#',
	}

	r := []rune(s)
	for i, c := range r {
		if v, ok := m[c]; ok {
			r[i] = v
		}
	}
	s = string(r)
	return s
}

func main13() {
	clear()
	welcometext := "----Welcome to Map Gen----"
	fmt.Println(welcometext)
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	for {
		var input string
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			fmt.Println("Failed to get Term Size")
			time.Sleep(time.Second * 2)
			clear()
		}
		fmt.Print(" [0] Generate\n [1] Read\n [2] Exit\n")
		fmt.Print(config.Prompt + " ")
		fmt.Scan(&input)
		if input == "2" || input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}
		inputint, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
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
			clear()
			var filename string
			fmt.Println("Filename(without ending): ")
			fmt.Print(config.Prompt + " ")
			fmt.Scan(&filename)
			filename = filename + ".map"
			err = os.WriteFile(filename, []byte(mapped), 0644)
			if err != nil {
				panic(err)
			}
			fmt.Print("Saved!")
			time.Sleep(3 * time.Second)
			clear()
		} else if inputint == 1 {
			var filename string
			fmt.Println("Filename(without ending): ")
			fmt.Print(config.Prompt + " ")
			fmt.Scan(&filename)
			filename = filename + ".map"
			data, err := os.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			content := string(data)
			lines := strings.SplitN(content, "\n", 2)
			if len(lines) < 2 {
				panic("Invalid map file")
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
				clear()
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

func main14() {
	green := "\033[32m"
	reset := "\033[0m"
	clear()
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting Terminal size:", err)
		return
	}
	for {
		softclear()
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

func typeeffect(ctx context.Context, input string, name string, text string) {
	string2 := ""
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		_, height = 80, 24
	}

	for i := 0; i < len(input); i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		clear()
		fmt.Println(text)
		string2 += string(input[i])
		fmt.Println("-----" + name + "-----")
		fmt.Println(string2)
		footerLine := height - 2
		promptLine := height - 1
		fmt.Printf("\033[%d;1H[Press 'q' or type 'exit' to exit]", footerLine)
		fmt.Printf("\033[%d;1H%s ", promptLine, config.Prompt)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 50):
		}
	}
}

func main15() {
	clear()
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
			clear()
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

func main16() {
	welcometext := "----Welcome to Sysmonitor----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	exitChan := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				exitChan <- true
				return
			}
		}
	}()
	for {
		select {
		case <-exitChan:
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			return
		default:
			clear()
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				width = 80
			}
			fmt.Println(welcometext)
			percentages, _ := cpu.Percent(0, false)
			cpuHeader := "-----CPU-----"
			cpuText := fmt.Sprintf("%.2f%%", percentages[0])
			cpuBar := fmt.Sprintf("[%s%s]",
				strings.Repeat("#", int(percentages[0])/5),
				strings.Repeat(" ", 20-int(percentages[0]/5)),
			)
			fmt.Println(strings.Repeat(" ", (width-len(cpuHeader))/2) + cpuHeader)
			fmt.Println(strings.Repeat(" ", (width-len(cpuText))/2) + cpuText)
			fmt.Println(strings.Repeat(" ", (width-len(cpuBar))/2) + cpuBar)
			v, _ := mem.VirtualMemory()
			ramHeader := "-----RAM-----"
			ramText := fmt.Sprintf("%.2f%% (%.0f/%.0f MB)",
				v.UsedPercent, float64(v.Used)/1024/1024, float64(v.Total)/1024/1024)
			ramBar := fmt.Sprintf("[%s%s]",
				strings.Repeat("#", int(v.UsedPercent)/5),
				strings.Repeat(" ", 20-int(v.UsedPercent/5)),
			)
			fmt.Println(strings.Repeat(" ", (width-len(ramHeader))/2) + ramHeader)
			fmt.Println(strings.Repeat(" ", (width-len(ramText))/2) + ramText)
			fmt.Println(strings.Repeat(" ", (width-len(ramBar))/2) + ramBar)
			fmt.Println("[Press 'q' or type 'exit' to exit]")
			fmt.Printf("%s ", config.Prompt)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func printNumber2(s string, replacement string) {
	lines := make([][]string, len(s))
	maxLines := 0

	for i, c := range s {
		digitStr := strings.ReplaceAll(digits2[c], "#", replacement)
		lines[i] = strings.Split(strings.Trim(digitStr, "\n"), "\n")

		if len(lines[i]) > maxLines {
			maxLines = len(lines[i])
		}
	}

	for i := 0; i < maxLines; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				fmt.Print(lines[j][i], "  ")
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

func main17() {
	symbol := "█"
	clear()
	welcometext := "----Welcome to ClockV2----"
	minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
	fmt.Println(welcometext)
	time.Sleep(time.Second)
	exitChan := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				exitChan <- true
				return
			}
		}
	}()
	for {
		select {
		case <-exitChan:
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			return
		default:
			clear()
			now := time.Now()
			Time := now.Format("15:04:05")
			printNumber2(Time, symbol)
			fmt.Println("[Press 'q' or type 'exit' to exit]")
			fmt.Printf("%s ", config.Prompt)
			time.Sleep(time.Second)
		}
	}
}

func main18() {
	animationmode := 0
	exitChan := make(chan bool)
	fmt.Println("[You can Type 'q' to exit]")
	time.Sleep(time.Second * 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "q" || input == "exit" {
				exitChan <- true
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-exitChan:
				return
			default:
				time.Sleep(time.Duration(rand.Intn(45-25)+25) * time.Second)
				if animationmode < 6 {
					animationmode++
				} else {
					animationmode = 0
				}
			}
		}
	}()
	for {
		select {
		case <-exitChan:
			clear()
			fmt.Print("---Bye!---")
			return
		default:
			if animationmode == 0 {
				width, _, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					width = 80
				}
				rows := []string{}
				delay := 40 * time.Millisecond
				for animationmode == 0 {
					row := ""
					for i := 0; i < width; i++ {
						row += "#"
						select {
						case <-exitChan:
							return
						default:
						}
						clear()
						for _, r := range rows {
							fmt.Println(r)
						}
						fmt.Print(row)
						time.Sleep(delay)
					}
					rows = append(rows, row)
					clear()
					for _, r := range rows {
						fmt.Println(r)
						select {
						case <-exitChan:
							return
						default:
						}
					}
					dropRow := strings.Repeat(" ", width-1) + "#"
					fmt.Print(dropRow)
					time.Sleep(delay)
					rows = append(rows, dropRow)
					row = ""
					for i := 0; i < width; i++ {
						select {
						case <-exitChan:
							return
						default:
						}
						row = "#" + row
						paddedRow := strings.Repeat(" ", width-len(row)) + row
						clear()
						for _, r := range rows {
							fmt.Println(r)
						}
						fmt.Print(paddedRow)
						time.Sleep(delay)
					}
					rows = append(rows, row)
					clear()
					for _, r := range rows {
						fmt.Println(r)
						select {
						case <-exitChan:
							return
						default:
						}
					}
					dropRow = "#" + strings.Repeat(" ", width-1)
					fmt.Print(dropRow)
					time.Sleep(delay)
					rows = append(rows, dropRow)
					_, height, err := term.GetSize(int(os.Stdout.Fd()))
					if err != nil {
						height = 24
					}
					if len(rows) >= height-1 {
						rows = []string{}
					}
				}
			} else if animationmode == 1 {
				symbol := "█"
				for animationmode == 1 {
					clear()
					select {
					case <-exitChan:
						return
					default:
					}
					now := time.Now()
					Time := now.Format("15:04:05")
					printNumber2(Time, symbol)
					time.Sleep(time.Second)
				}
			} else if animationmode == 2 {
				width, _, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					panic(err)
				}
				var length int
				var mode int
				mode = 0
				length = 10
				for animationmode == 2 {
					select {
					case <-exitChan:
						return
					default:
					}
					if mode == 0 {
						if length == width || length > width {
							mode = 1
						} else {
							if length+1 == width {
								length += 1
							} else {
								length += 2
							}
						}
					} else if mode == 1 {
						length--
						if length == 10 {
							mode = 0
						}
					}
					clear()
					fmt.Print((strings.Repeat(strings.Repeat(" ", (width/2-length/2))+strings.Repeat("#", length)+"\n", width-1)))
					time.Sleep(time.Millisecond * 50)
				}
			} else if animationmode == 3 {
				clear()
				width, _, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					panic(err)
				}
				heights := []int{0, 0, 0, 0, 0}
				for animationmode == 3 {
					select {
					case <-exitChan:
						return
					default:
					}
					for i := 0; i < len(heights); i++ {
						select {
						case <-exitChan:
							return
						default:
						}
						if heights[i] > width {
							heights[i] = 0
						}
						heights[i] += 1
						fmt.Println(strings.Repeat("#", heights[i]) + "#")
						time.Sleep(time.Millisecond * 200)
					}
				}
			} else if animationmode == 4 {
				width, height, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					panic(err)
				}

				heights := make([]int, height)
				for animationmode == 4 {
					select {
					case <-exitChan:
						return
					default:
					}
					for i := 0; i < len(heights); i++ {
						select {
						case <-exitChan:
							return
						default:
						}
						clear()
						if heights[i] > width {
							heights[i] = 0
						}
						heights[i] += 1
						for i2 := 0; i2 < len(heights); i2++ {
							fmt.Println(strings.Repeat(" ", heights[i2]) + "#")
						}
						time.Sleep(time.Millisecond * 50)
					}
				}
			} else if animationmode == 5 {
				green := "\033[32m"
				reset := "\033[0m"
				clear()
				width, height, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					fmt.Println("Error getting Terminal size:", err)
					return
				}
				for animationmode == 5 {
					select {
					case <-exitChan:
						return
					default:
					}
					softclear()
					string := ""
					for i := 0; i < width*height; i++ {
						select {
						case <-exitChan:
							return
						default:
						}
						if rand.Intn(2) == 0 {
							string += "0"
						} else {
							string += "1"
						}
					}
					fmt.Print(green + string + reset)
					time.Sleep(100 * time.Millisecond)
				}
			} else if animationmode == 6 {
				width, height, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					fmt.Println("Error getting Terminal size:", err)
					return
				}
				heights := []int{0, 1, 2, 3, 2, 1, 0}
				frame := 0
				lineLength := height

				for animationmode == 6 {
					select {
					case <-exitChan:
						return
					default:
					}
					clear()
					for i := 0; i < lineLength; i++ {
						select {
						case <-exitChan:
							return
						default:
						}
						y := heights[(frame+i)%len(heights)]
						fmt.Printf("%*s#\n", (width/2)+y, "")
					}
					frame++
					time.Sleep(time.Millisecond * 100)
				}
			}
		}
	}
}

func main19() {
	var SERVER string
	fmt.Println("Enter Server URL [default: wss://chat.bolucraft.uk/chat]")
	fmt.Print(config.Prompt + " ")
	fmt.Scanln(&SERVER)
	if SERVER == "" {
		SERVER = "wss://chat.bolucraft.uk/chat"
	}

	conn, _, err := websocket.DefaultDialer.Dial(SERVER, nil)
	if err != nil {
		fmt.Println("Error Connecting:", err)
		time.Sleep(time.Second * 2)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to ChatServer", SERVER)
	fmt.Println("[Type 'q' or 'exit' to exit]")

	done := make(chan bool)
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("\nDisconnected from server.")
				done <- true
				return
			}
			fmt.Println("\n" + string(message))
			fmt.Print(config.Prompt + " ")
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(config.Prompt + " ")
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		conn.WriteMessage(websocket.TextMessage, []byte(text))
		if strings.ToLower(text) == "exit" || strings.ToLower(text) == "q" {
			break
		}
	}

	<-done
	fmt.Println("Client Closed.")
}

// START Config
type Config struct {
	FirstRun bool   `json:"first_run"`
	Prompt   string `json:"prompt"`
}

func configure() {
	configFile := "config.json"
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			config = Config{
				FirstRun: true,
				Prompt:   "",
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Welcome! Please set your Prompt(Default: '>'): ")
			input, _ := reader.ReadString('\n')
			config.Prompt = strings.TrimSpace(input)
			config.FirstRun = false
			saveConfig(configFile, config)
			fmt.Println("Prompt Saved! Starting SimpleApps...")
		} else {
			fmt.Println("Error Reading Config:", err)
			return
		}
	} else {
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Println("Error Parsing Config:", err)
			return
		}
	}
}

func saveConfig(filename string, cfg Config) {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("Error Marshalling Config:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error Saving Config:", err)
	}
}

//END Config

func menu() {
	rand.Seed(time.Now().UnixNano())
	configure()
	data, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println("Error reading local version:", err)
		time.Sleep(2 * time.Second)
		return
	}
	localVersion := strings.TrimSpace(string(data))
	for {
		clear()
		welcometext := "----Welcome to SimpleApps----"
		minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
		programs := []func(){main1, main2, main3, main4, main5, main6, main7, main8, main9, main10, main11, main12, main13, main14, main15, main16, main17, main18, main19, update}
		names := []string{"NumberChecker", "GradeChecker", "UnitConverter", "Number2Bar", "CoinFlip", "Countdown", "Timer", "Clock", "Magic 8-Ball", "800+ Lines Special", "Calculator", "ToDo List", "Map Gen", "Matrix", "FakeLogGen", "SysMonitor", "ClockV2", "ASCII Animations", "SimpleChat", "Update"}
		fmt.Println(welcometext)
		fmt.Println(strings.Repeat("-", 6) + "You are on V" + localVersion + strings.Repeat("-", 6))
		fmt.Println("What do you want to run?")
		for i, name := range names {
			fmt.Printf("[%d] %s\n", i, name)
		}
		fmt.Printf("[%d] Exit\n", len(programs))
		fmt.Printf("%s ", config.Prompt)
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		}

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice, try again!")
			continue
		}

		if choice == len(programs) {
			fmt.Println(strings.Repeat("-", minuses), "Bye!", strings.Repeat("-", minuses))
			time.Sleep(3 * time.Second)
			clear()
			break
		} else if choice >= 0 && choice < len(programs) {
			programs[choice]()
		} else {
			fmt.Println("Invalid choice, try again!")
			clear()
		}
	}
}

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting Terminal size:", err)
		return
	}
	fmt.Printf("Terminal: %d columns, %d lines\n", width, height)

	minWidth := 70
	minHeight := 23

	if width < minWidth || height < minHeight {
		fmt.Printf("[Warning] Terminal is too small, min size: %dx%d\n", minWidth, minHeight)
	} else {
		menu()
	}
}
