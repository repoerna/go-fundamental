package other

import "fmt"

var privateVariable = "this is private"
var PublicVariable = "this is public"

func HelloFromOtherPackage(name string) {
  fmt.Println("Hello", name, ", this is from other package")
}