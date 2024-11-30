package main

import (
    "fmt"
    "math"
)

func split(num, precision float64) (integer, dec int) {
    integer = int(num)
    mul_val := math.Pow(10, precision)

    // Using round as there is too much precision
    dec = int(math.Round((num - float64(integer)) * mul_val))

    return  // values not required as specified in func def
}

func main() {
    fmt.Println(split(1.325, 3))
}
