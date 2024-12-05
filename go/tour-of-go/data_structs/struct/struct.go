package main

import "fmt"

// define a structure
type Vertex struct {
    X int
    Y int
}

func main() {
    a := Vertex{3, 4}
    fmt.Println(a)

    // access fields
    a.X = 4
    fmt.Println(a)  // {4, 4}

    // through pointers
    p := &a

    (*p).X = 3 // actual syntax
    p.X = 4 // go allows this (auto dereference)

    // Struct literals
    var (
        v1 = Vertex{1, 2}  // has type Vertex
        v2 = Vertex{X: 1}  // Y:0 is implicit
        v3 = Vertex{}      // X:0 and Y:0
        p  = &(Vertex{1, 2}) // has type *Vertex
    )
}
