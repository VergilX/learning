/*
THIS FILE IS NOT MEANT TO BE RUN

Use this for reviewing syntax
*/


package main    // name of current package

// Imports
import (
    "fmt"
    "math"
)

// or use this syntax
import "fmt"
import "math"


// Global var
var max, queue = 10, 20


// function declaration
func square(num int) (float64) {
    return math.Pow(int, 2)
}


// Named return func
func square_and_print(num int) (ans float64) {
    ans = math.Pow(int, 2) // not declared (check func return type)
    fmt.Println(ans)
    return
}


// main func
func main() {
    // Assignment
    var first, last string

    first = "Abhinand"
    last = "Manoj"

    // initializers don't require type
    var first, last = "Abhinand", "Manoj"


    // := used instead of var (not available outside func body)
    firstname := "Abhinand"
    lastname := "Manoj"

    // Default values
    var i int // 0
    var j bool // false
    var name string // ""

    // Type conversion
    f := float64(i)
    dec_part := (math.Pi - int(math.Pi))

    // Constants (don't use := syntax only const)
    const Pi = 3.14
    const (
        g = 9.8
        light = 3.8
    )
}
