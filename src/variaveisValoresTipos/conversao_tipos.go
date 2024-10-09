package main

import "fmt"

var a int = 20
var b string = "Glauco Santos"

func main() {
	var c float64 = float64(a) // CONVERSÃO -> tipo(variável)
	
	fmt.Printf("O valor de C é: %g e o valor de B é: %s ", c, b)
	
	// Descobrindo o tipo de uma variável
	fmt.Printf("O tipo de A é: %d (%T), o tipo de B é: %T e o tipo de C é: %T", a,a, b, c)

}