package apps

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

func Main18() {
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
			Clear()
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
						Clear()
						for _, r := range rows {
							fmt.Println(r)
						}
						fmt.Print(row)
						time.Sleep(delay)
					}
					rows = append(rows, row)
					Clear()
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
						Clear()
						for _, r := range rows {
							fmt.Println(r)
						}
						fmt.Print(paddedRow)
						time.Sleep(delay)
					}
					rows = append(rows, row)
					Clear()
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
					Clear()
					select {
					case <-exitChan:
						return
					default:
					}
					now := time.Now()
					Time := now.Format("15:04:05")
					PrintNumber2(Time, symbol)
					time.Sleep(time.Second)
				}
			} else if animationmode == 2 {
				width, _, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					fmt.Println("Error:", err)
					return
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
					Clear()
					fmt.Print((strings.Repeat(strings.Repeat(" ", (width/2-length/2))+strings.Repeat("#", length)+"\n", width-1)))
					time.Sleep(time.Millisecond * 50)
				}
			} else if animationmode == 3 {
				Clear()
				width, _, err := term.GetSize(int(os.Stdout.Fd()))
				if err != nil {
					fmt.Println("Error:", err)
					return
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
					fmt.Println("Error:", err)
					return
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
						Clear()
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
				Clear()
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
					Softclear()
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
					Clear()
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
