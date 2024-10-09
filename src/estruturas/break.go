package main

import "fmt"

func main() {

	// Loop infinito com ponto de parada
	cont:=0
	for  {
		if cont <=20 {
			fmt.Println("contador <= 20")
			cont++
		}else{
			fmt.Println("contador é maior 20")
			break
		}
	}

	fmt.Println()
	
	for i := 0; i < 20; i++ {
		if  (i>=3 && i<=5) || (i>=10 && i<=15){
			continue // encaminha a execução para o final o loop
		}
		fmt.Printf("%d ", i)
	}

}