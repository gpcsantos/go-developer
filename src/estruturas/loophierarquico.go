package main

import "fmt"

func main(){
	for hora := 0; hora <=12; hora++ {
		fmt.Print("Hora: ", hora , " - Minutos: ")
		for min := 1; min < 60; min++ {
			fmt.Printf("%d ",min)
		}
		fmt.Print("\n") // quebra de linha  OU fmt.Println()
	}
}
// gobyexample.com