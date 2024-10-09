package main

import "fmt"

var x string // escopo Global

func main() {
	x = "Hello World"

	fmt.Println(x)

	// x= 1  // vai dar erro 
	// fmt.Println(x)

}