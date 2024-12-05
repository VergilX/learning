package main

import "fmt"


var MAX = 1024
var buffer [1024]int // cannot be resized

func main() {
    // ********************************** DECLARED VARS *********************************
    var vowels []string  // no size specified
    vowels = []string{"a", "e", "i", "o", "u"} // size infered
    // vowels = [...]string{"a", "e", "i", "o", "u"}  // invalid statement



    var prime [5]int  // size specified
    prime = [...]int{1, 2, 3, 5, 7}  // ... required for size inference
    // prime = []int{1, 2, 3, 5, 7}  // invalid statement (... required)


    // SHORTHAND VERSION (undeclared)
    my_name := [3]string{"Abhinand", "D", "Manoj"}
    names := []string{"Abhinand", "Nimil"}  // inferred size (... can be put)
    consonants := [6]string{"b", "c", "d", "f", "j"}  // basically a slice (< 6)

    // CONCLUSION: ... is used for inference after size mentioned in declaration
    // IMPORTANT: Not using a length actually creates a slice and returns it


    // **********************************************************************************
    // Slices: dynamically-sized and flexible
    // Modifications in slices are referenced in originals
    // Conventions similar to python
    // Info: https://go.dev/blog/slices-intro

    // create a slice using previously defined vowels array
    slice := vowels[ : len(vowels)]  // [e i o u]
    var slice2 []string = vowels[1 : 4]    // [e i o]

    // modify slice (reflected in original)
    slice[2] = "o"
    fmt.Println("New:", vowels)   // New: [a e o o u]

    // Slice literal: Returns a slice of the array
    array := [5]int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3, 4, 5}

    // Slice: length and capacity
    array = [5]int{1, 2, 3, 4, 5}
    slice = array[1:4]

    fmt.Println(len(slice)) // 3  (currently used)
    fmt.Println(cap(slice)) // 5  (total available memory)

    // zero value of a slice is nil
    var s []int  // nil; len(s)=nil; cap(s)=nil

    // creating slice with make
    a := make([]int, 5)   // make([]int, len(), cap())

    // ***********************************************************************************
    // slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

    // Appending a slice
    part1 := []int{1, 2, 3, 4, 5}
    part2 := []int{6, 7, 8, 9, 10}
    sum1 := append(part1, 6, 7, 8, 9, 10)
    sum2 := append(part1, part2...)

    // info: If the backing array of slice is too small
    // a bigger array will be allocated

    // ************************************ RANGE *****************************************************

    // Range
    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
    // i: index, v: copy of elem at index

    // to avoid one of them
    for _, v := range pow {
        // do something
        fmt.Println(v)
    }


    fmt.Println(&slice == &vowels)  // interesting idk the implementation

    /*
    // Uncomment this section to run this code (ignore output)
    fmt.Println(vowels, prime, names, consonants, my_name)
    fmt.Println(slice2)
    */

}
