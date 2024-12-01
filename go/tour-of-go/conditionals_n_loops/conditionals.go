package main

import "fmt"

func main() {
    x := -99

    // no omittance of {} even for single line body (unlike C)
    if x < 0 {
        fmt.Println("negative")
    }

    // init var (scope in if block only)
    firstname := "Abhinand"

    if lastname := "Manoj"; firstname == "Abhinand" {
        fmt.Println("You're awesome")
    } else {
        fmt.Println("You're gay")
    }
    fmt.Println(lastname)  // error

    daytime := "Noon"
    // switch-case
    switch day {
        case "Morning":
            fmt.Println("Good Morning")
        case "Afternoon":
            fmt.Println("Good Afternoon")
        case "Evening":
            fmt.Println("Good Evening")
        case "Night":
            fmt.Println("Good Night")
        default:
            fmt.Println("Have a good day!")
    }

    // switch with no condition ~ switch true {}
    switch {
        case daytime == "Morning":
            fmt.Println("Good Morning")
        case daytime == "Afternoon":
            fmt.Println("Good afternoon")
        default:
            fmt.Println("Have a good day!")
    }
}
