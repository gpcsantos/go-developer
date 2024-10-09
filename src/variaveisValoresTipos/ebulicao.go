package main

import "fmt"

// variávies estaticamente tipadas
const ebulicaoF float64 = 212.0

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println("A tempenratura de ebulição da água em ºF = ", tempF)
	fmt.Println("A tempenratura de ebulição da água em ºC = ", tempC)

}