package apps

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, indent string, recursive bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var entryhidden bool
	for _, entry := range entries {
		var extrasymbol string
		if entry.IsDir() {
			extrasymbol = "/"
		} else {
			extrasymbol = ""
		}
		if strings.HasPrefix(entry.Name(), ".") {
			entryhidden = true
		} else {
			entryhidden = false
		}
		if entryhidden == true && recursive == true {
			fmt.Println(indent + entry.Name() + extrasymbol)
			if entry.IsDir() {
				printTree(filepath.Join(path, entry.Name()), indent+"  ", recursive)
			}
		} else {
			if entryhidden == false {
				fmt.Println(indent + entry.Name() + extrasymbol)
				if entry.IsDir() {
					printTree(filepath.Join(path, entry.Name()), indent+"  ", recursive)
				}
			}
		}
	}
}

func Main26() {
	Greet("Tree")
	for {
		var recursivebool bool
		dir := Input("What dir(or type ':exit' to exit)?")
		if dir == ":exit" {
			Bye()
			return
		}
		for {
			recursive := Input("Recursive(also hidden files)?")
			if recursive == "y" || recursive == "Y" {
				recursivebool = true
				break
			} else if recursive == "n" || recursive == "Y" {
				recursivebool = false
				break
			} else {
				continue
			}
		}
		printTree(dir, "", recursivebool)
	}
}
