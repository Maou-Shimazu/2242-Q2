package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        loop(12)
        wg.Done()
    }()

    go func() {
        loop(8)
        wg.Done()
    }()

    wg.Wait()
}

func loop(times int) {
    for i := 0; i < times; i++ {
        fmt.Println("Running", i, "times")
        time.Sleep(time.Millisecond * 500)
    }
}
