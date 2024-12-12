package main

import "fmt"

// declare struct
type Saiyan struct {
    Name string
    power int
}


func main() {
    println("It's over 9000!")

    // init struct
    goku := Saiyan {
        Name: "Goku",
        power: 9000,
    }

    vegeta := Saiyan {
        "Vegeta",
        8000,
    }

    gohan := new(Saiyan)
    // same as
    gohan := &Saiyan{}

    fmt.Println(goku, vegeta)
}
