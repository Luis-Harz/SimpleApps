package apps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/term"
)

func Main16() {
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
			Clear()
			return
		default:
			Clear()
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
			fmt.Printf("%s ", ConfigData.Prompt)
			time.Sleep(time.Millisecond * 300)
		}
	}
}
