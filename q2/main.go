package main

import (
  "fmt"
)

func main() {
  channelMultiplication := make(chan int)

  go func() { channelMultiplication <- 20*5 }()

  result := <- channelMultiplication
  fmt.Println(result)
}
