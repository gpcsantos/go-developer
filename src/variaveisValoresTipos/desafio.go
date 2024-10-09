package main

import "fmt"

var c float64
var ebulicaok int = 373

func main() {


	c = float64(ebulicaok - 273)

	fmt.Printf("O ponto de ebulição da água é %d em Kelvin que corresponde à %g° Celsius", ebulicaok, c)
}