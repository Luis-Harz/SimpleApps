package apps

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

	"golang.org/x/term"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadInput() string {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

var filetodownload string = ""
var serverurl string = "simplemirror.bolucraft.uk/ForServer"

// ASCII Digits START
var Digits = map[rune]string{
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
var Digits2 = map[rune]string{
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
func Random(minimum int, maximum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(maximum-minimum+1) + minimum
	return n
}

func Clear() {
	fmt.Print("\033[H\033[2J")
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
		Clear()
		fmt.Println(text)
		string2 += string(input[i])
		fmt.Println("-----" + name + "-----")
		fmt.Println(string2)
		footerLine := height - 2
		promptLine := height - 1
		fmt.Printf("\033[%d;1H[Press 'q' or type 'exit' to exit]", footerLine)
		fmt.Printf("\033[%d;1H%s ", promptLine, ConfigData.Prompt)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 50):
		}
	}
}

// Help Functions main13
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

//END Help Functions main13

// ====================Help Functions main12====================
type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

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
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(todos); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

//END Help Functions main12

//Help Functions main11
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

//END Help Functions main11

// Help Functions main10
func randChar() byte {
	if rand.Intn(2) == 0 {
		return '#'
	}
	return ' '
}

//END Help Functions main10

// Help Function main8
func printNumber(s string) {
	lines := make([][]string, len(s))
	maxLines := 0
	for i, c := range s {
		lines[i] = strings.Split(strings.Trim(Digits[c], "\n"), "\n")
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

//END Help Function main8

// Help Function main7
func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

//END Help Function main7

func PrintNumber2(s string, replacement string) {
	lines := make([][]string, len(s))
	maxLines := 0

	for i, c := range s {
		digitStr := strings.ReplaceAll(Digits2[c], "#", replacement)
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

func Softclear() {
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
	if filename == "SimpleApps" || filename == "Simpleexe" {
		os.Chmod(filename, 0755)
	}
	return err
}

func Update() {
	Clear()
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
	fmt.Print(ConfigData.Prompt + " ")
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
		fileName = "Simpleexe"
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
