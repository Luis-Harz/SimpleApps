package apps

import (
	"fmt"
	"time"
)

func Main24() {
	for {
		Clear()
		Greet("Wget")
		url := Input("URL(or type 'exit' to exit): ")
		if url == "exit" {
			Bye()
			return
		}
		location := Input("Location: ")
		downloadFile(url, location)
		fmt.Println("Saved file sucessfully!")
		time.Sleep(time.Second * 2)
	}
}
