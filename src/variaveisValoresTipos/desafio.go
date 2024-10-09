package main

import "fmt"

var k float64
var ebulicaoC int = 100

func main() {


	k = float64(ebulicaoC + 273)

	fmt.Printf("O ponto de ebulição da água é %d° Celsius que corresponde à %g Kelvin", ebulicaoC, k)
}