package apps

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Main23() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to start")
	for {
		reader.ReadString('\n')
		fmt.Print("\033[2J\033[H")
		resp, err := http.Post("https://simpleclick.bolucraft.uk/click", "text/plain", nil)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("New count:", string(body))
	}
}
