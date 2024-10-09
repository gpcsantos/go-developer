package main

import "fmt"

const ebulicaoF float64 = 212.0
// ebulicaoF := 212.0 // irá apresentar um erro, pois criar variável por atriução precisar ser em um bloco

func main() {
	tempF := ebulicaoF  // cria variável atribui valor por inferencia, precisa estar em um bloco
	tempC := (tempF - 32) * 5 / 9 // cria variável atribui valor por inferencia

	fmt.Println("A tempenratura de ebulição da água em ºF = ", tempF)
	fmt.Println("A tempenratura de ebulição da água em ºC = ", tempC)

}