package main

import (
    "fmt"
    "time"
)

// normal defer (prints hello world)
func hello() {
    defer fmt.Println("world")

    fmt.Println("hello, ")
}

// stacked defer (implements a stopwatch 0-sec)
func stopwatch(sec int) {
    for ;sec >= 0; sec-- {
        defer fmt.Println("Timer:", sec)
        defer time.Sleep(time.Second)
    }
}

func main() {
    // call anything
    hello()
    stopwatch(10)
}
