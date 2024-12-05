// Question: https://go.dev/tour/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    start := false
    fib := 0
    a := 0
    b := 1

    // closure function
    return func() int {
        if fib == 0 {
            fib = 1
            return 0
        } else if fib == 1 && !start  {
            fib = 1
            start = true
        } else {
            fib = a + b
            a = b
            b = fib
        }

        return fib
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
