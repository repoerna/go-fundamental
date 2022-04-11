package main

import "fmt"

func main() {
  fmt.Println(`this is back quote string`)
  fmt.Println("this is double quote string")

  // built in function in string type
  fmt.Println(len("hello")) // count characters
  fmt.Println("hello"[1]) // get characters using index position, will return raw byte

}

 // https://gobyexample.com/strings-and-runes
