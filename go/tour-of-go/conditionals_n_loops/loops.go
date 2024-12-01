package main

import "fmt"

func main() {
    /*
    for loop format (used for while too):
        - initialize
        - condition
        - update

    Only condition is compulsary
    No paranthesis
    */

    // i is only available in loop scope
    for i := 10; i >= 0; i-- {
        fmt.Println(i)
    }
    fmt.Println(i) // error

    var i int = 10
    for ;i >= 0; i-- { // notice the ;
        fmt.Println(i)
    }

    // while loop
    i = 10
    for i >= 0 {
        fmt.Println(i)
        i--
    }

    // infinite loop
    for {
        fmt.Println("nimil is definitely gay")  // easter egg
    }
}
