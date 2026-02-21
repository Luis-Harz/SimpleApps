package apps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// START main12

// ====================Main12====================

func Main12() {
	Clear()
	welcometext := "----Welcome to ToDo List----"

	for {
		fmt.Println(welcometext)
		minuses := (len(welcometext) / 2) - (len(" Bye! ") / 2)
		scanner := bufio.NewScanner(os.Stdin)
		listTodos()
		fmt.Printf("%s ", ConfigData.Prompt)
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
			Clear()
			break
		} else {
			fmt.Println("Unknown Command. Type 'help'")
		}
		time.Sleep(time.Second * 2)
		Clear()
	}
}

//END Main12
