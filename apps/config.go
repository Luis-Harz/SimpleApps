package apps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var ConfigData Config

// START Config
type Config struct {
	FirstRun bool   `json:"first_run"`
	Prompt   string `json:"prompt"`
}

func version() string {
	data, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println("Error reading local version:", err)
		time.Sleep(2 * time.Second)
	}
	localVersion := strings.TrimSpace(string(data))
	return localVersion
}

func Configure() {
	configFile := "config.json"
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			ConfigData = Config{
				FirstRun: true,
				Prompt:   "",
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Welcome! Please set your Prompt(Default: '>'): ")
			input, _ := reader.ReadString('\n')
			ConfigData.Prompt = strings.TrimSpace(input)
			ConfigData.FirstRun = false
			saveConfig(configFile, ConfigData)
			fmt.Println("Prompt Saved! Starting SimpleApps...")
		} else {
			fmt.Println("Error Reading Config:", err)
			return
		}
	} else {
		err = json.Unmarshal(data, &ConfigData)
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
