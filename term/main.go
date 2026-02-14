package main

import (
   "fmt"
)

func main() {
test := 0
for {
   if test < 10000000 {
      test++
   } else {
      fmt.Println("Reached!")
      break
   }
}
}
