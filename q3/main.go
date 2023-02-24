package main

import (
  "fmt"
)

func main() {
  fmt.Println("Starting")
  go func(times int) {
    for i := 0; i < times; i++ {
      fmt.Println("Running", i, "now.")
    }
  }(4)
  fmt.Println("Done")
  fmt.Scanln()
}


