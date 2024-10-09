package main

import "fmt"

func main() {
	fmt.Println(len(" Hello World"))
	fmt.Println("Hello World" [2]) // mostra 108 é "l" Decimal na tabela ASCII
	fmt.Println("Hello" + "World") // não dá epaço
	fmt.Println("Hello" , "World") // adiciona o espaço entre as palavras
	fmt.Println(`Hello World com crase...`)
}