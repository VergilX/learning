package main

import "fmt"

type Person struct {
    SSNum int
    Name string
    Address string
}

// No need for same func defined in parent
type Employee struct {
    *Person    // "inherits" Person attr
    EmpID int
    Salary int // won't be needing int :[
}

func main() {
    me := &Employee {
        Person: &Person {
            SSNum: 1001,
            Name: "Rodrick",
            Address: "New Jersey",
        },

        EmpID: 200222,
        Salary: 999999, // trailing `,` is required
    }


    fmt.Println(me.Person.Name)  // Rodrick
    fmt.Println(me.Name)         // Rodrick

}
