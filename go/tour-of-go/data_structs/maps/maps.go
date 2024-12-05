package main

import "fmt"

// Global declaration
var m map[string]int  // create a map m, mapping from string(key) to int(value)

type Vertex struct  {
    X, Y float64
}

func main() {
    // using make
    logged_in := make(map[string]bool)  // maps from string to string

    // zero value of map is nil

    // add value
    logged_in["Abhinand"] = true

    // I didn't get map literals
    var location = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},
        "Google":    {37.42202, -122.08408},
    }

    // delete element
    delete(location, "Google")

    // test if elem exists
    elem, ok := location["Google"]  // ok is bool value
    fmt.Println(elem, ok) // 0 false


    fmt.Println(m, logged_in["Abhinand"], location)
}
