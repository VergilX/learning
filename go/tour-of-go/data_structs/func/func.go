package main

import "fmt"

// can be used to print result of function fn
func printResult(fn func(int, int) (int), a int, b int) {
    res := fn(a, b)
    fmt.Println(res)
}

func main() {
    // adder function
    adder := func (a, b int) (int) {
        return a+b
    }

    // multiplier function
    mul := func (a, b int) (int) {
        return a*b
    }

    a := 10
    b := 20
    printResult(adder, a, b)
    printResult(mul, a, b)


    // Function closures: Functions which reference vars
    // outside it's scope
    // not writing an example, refer:
    // https://go.dev/tour/moretypes/25
}
