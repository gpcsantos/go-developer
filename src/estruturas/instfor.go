package main

import "fmt"

func main(){
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")
	fmt.Println("")

	cont := 1
	for cont <=10{
		fmt.Println(cont)
		cont = cont + 1
	}
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}