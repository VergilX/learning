package main

import (
    "fmt"
    "math"
)

// type defined in same file
type Vertex struct {
    X, Y int
}

// Methods are defined on receiver argument
// receiver type should be defined in same file

// method on Vertex (defined b/w func and name of func)
func (v Vertex) Scale(factor int) Vertex {
    var vertex Vertex

    vertex.X = v.X * factor
    vertex.Y = v.Y * factor

    return vertex
}
// method = func with receiver arg

// receiver can be non struct
type MyInt int

// using pointer receivers, pass by reference
func (i *MyInt) Square() {
    i = &MyInt(math.Pow(float64(*i), 2))
}

func main() {
    v := Vertex{3, 4}
    newv := v.Scale(10)
    fmt.Println(newv)

    num := MyInt(10)
    num.Square()
    fmt.Println(num)
}
