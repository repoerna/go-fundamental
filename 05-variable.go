package main

import "fmt"

func main() {

    var firstVar string // declaration
    firstVar = "this will initiate the variable" // initialization
    fmt.Println(firstVar)
    

  // declare multiple variables at once with same type.
    var b, c int = 1, 2
    fmt.Println(b, c)

  // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0 and fro string is "".
    var e int
    fmt.Println(e)

  // Go will infer the type of initialized variables.
    var a = "initial"
    fmt.Println(a)

    var d = true
    fmt.Println(d)

  // clean version
    var (
      aString string = "this is string"
      aInt = 1
    )
    fmt.Println(aString, aInt)

  // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
    f := "apple"
    fmt.Println(f)
}