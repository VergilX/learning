package main

import "fmt"

func main() {
    // Declaration
    var p *int

    fmt.Println(p) // <nil>

    // & operator
    i := 42
    p = &i  // p contains address of i

    // Dereference
    fmt.Println(*p) // access i through p
    *p = 45  // set i through p

    fmt.Println(i) // 45


    // Go has NO POINTER ARITHMETIC

}
