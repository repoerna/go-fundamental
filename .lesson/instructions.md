# Go Fundamental
## Overview
### Go History
- Created at Google in 2007
- Release on public on 2009
- Became popular when Docker built using Go in 2011

### Why Go
- Go is a [compiled](https://www.guru99.com/difference-compiler-vs-interpreter.html#:~:text=Compiler%20transforms%20code%20written%20in,while%20interpreted%20code%20runs%20slower.), concurrent, garbage-collected, statically typed language
- Go is efficient, scalable, and productive. 


### Understand Go Lifecycle/Development Process
![go lifecycle](assets/go-lifecycle.draw)

### More About Go and Other Resources
- [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)
- [gobyexample](https://gobyexample.com/)

### Install Go in local
If you want to install go in local you can read the [documentation](https://go.dev/doc/install) 

### Go Command
```sh
go <command> [arguments]
```
Most commonly used command:
- go run
  <br> compile and run Go program
- go build
  <br> to compile packages and dependencies
- go test
  <br> test packages
- go get
  </br> add dependencies to current module and install them

to see all available commands run this on console
```sh
go
```
to see detail of a command
```sh
go help <command>
```

## Go Basic
### Hello World (1)
#### Go Program Structure
In `1-hello-world.go` file you can see the code contains some parts:
- package
  <br> A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
- import
  <br> to import and use othe packages
- function main
  <br> main() function is a special type of function and it is the entry point of the executable programs. It does not take any argument nor return anything. 


### Number (2)
there are 2 basic type of numbers:
- integer
- floating point

#### Integer (int)
|data type|min value|max value|
|-|-|-|
|int8|-128|127|
|int16|-32768|32767|
|int32|-2147483648|2147483647|
|int864|-9223372036854775808|9223372036854775807|


#### Unsigned Integer (uint)
data type|min value|max value|
|-|-|-|
|uint8|0|255|
|uint16|0|65535|
|uint32|0|4294967295|
|uint864|0|18446744073709551615|

#### Floating Point (float)
data type|min value|max value|
|-|-|-|
|float32|1.18e−38|3.4e38|
|float64|2.23e-308|1.80e308|
|complex32|0|4294967295|
|complex64|0|18446744073709551615|

### Other Number
|data type|equal with|
|-|-|
|byte|uint8|
|rune|int32|
|int|min int32|
|uint|min uint32|

## Boolean (3)
Simple, boolean (bool) in Go just have 2 value
- true
- false

## String and Runes (4)
Strings are equivalent to `[]byte`
<br>
Go treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a `rune` - it’s an integer that represents a Unicode code point

[more about string](https://go.dev/blog/strings)

## Variable (5)
In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
To put is simple, variable is container to store the data.

To create variable in Go we can use keyword `var` and using `:=` syntax

## Constant (6)
Go supports constants of character, string, boolean, and numeric values.
Constant value can't be changed.
To create constant in Go, we can use keyword `const`.

## Arrays (7)
In Go, an array is a numbered sequence of elements of a **specific length**.

## Slice (8)
Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

## Maps (9)
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

## For (10)
`for` is Go’s only looping construct. Here are some basic types of `for` loops.
The syntax:

- infinite loop
```
for { }
```
- with only condition
```
for <condition> { } 
```
- with initial, condition, and post statement
```
for <initial statement> ; <condition> ; <post statement> { }
```

## If Else (11)
Branching with `if` and `else` in Go is straight-forward.
The syntax:
```
if <condition> { 
  ...
} else if <condition> {
  ...
} else {
  ...
}
```

## Range (12)
`range` iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.

## Functions (13)
Functions are central in Go. Function syntax is:
```
func <function-name> (<argument1> <argument-type1>, <argument2> <argument-type2>) (output-type1, output-type2) { }
```

## Variadic Function (14)
[Variadic functions](https://en.wikipedia.org/wiki/Variadic_function) can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function. The syntax:
```
...
```

## Closures (15)
Go supports [anonymous functions](https://en.wikipedia.org/wiki/Anonymous_function), which can form [closures](https://en.wikipedia.org/wiki/Closure_(computer_programming)). Anonymous functions are useful when you want to define a function inline without having to name it.

## Recursion (16)
Go supports [recursive functions](https://en.wikipedia.org/wiki/Recursion_(computer_science)).
Recursion is function that calls it self.
![rec](https://media.geeksforgeeks.org/wp-content/cdn-uploads/recursion.jpg)

### Struct (17)
Go’s structs are typed collections of fields. They’re useful for grouping data together to form records. Its like class on OOP Programming, but different.
The syntax
```
type <struct-name> struct {
  <field1> <field1-type>
  <field2> <field2-type>
  ...
}
```

## Method (18)
Methods is function that defined on struct type.
The syntax
```
func (<abbr> <struct-name>) <methd-name> <return-type> {
  ...
}
```

[Method Receiver - Pointer v/s Value](https://medium.com/globant/go-method-receiver-pointer-vs-value-ffc5ab7acdb)

## Interface (19)
Interfaces are named collections of method signatures.

## Pointer (20)
Go supports [pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming)), allowing you to pass references to values and records within your program.

![point](https://media.geeksforgeeks.org/wp-content/uploads/20190705160332/Pointers-in-Golang.jpg)

Pointer `a` pointing to the memory address associated with variable `b`. In this diagram, the computing architecture uses the same address space and data primitive for both pointers and non-pointers; this need should not be the case.

## Goroutine (21)
A goroutine is a lightweight thread of execution.

## Channel (22)
Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
![chan](https://media.geeksforgeeks.org/wp-content/uploads/20190731140438/Untitled-Diagram46.jpg)

## Defer (23)
Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.