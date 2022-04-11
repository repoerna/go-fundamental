package main

import (
	"fmt"
	"gofundamental/other"
)

func main() {
  other.HelloFromOtherPackage("test")

  fmt.Println(other.PublicVariable)

  fmt.Println(other.privateVariable) // will be error
}