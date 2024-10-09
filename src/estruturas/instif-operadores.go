package main

import "fmt"

func main(){
	x:=2
	if x==2 || x==3{
		fmt.Println("Sim, x é igual a 2 ou 3")
	}

	fmt.Println("Multimplos")
	m:= 21564665467

	if m%2==0 && m%3==0{
		fmt.Println(m,"é multiplo de 2 e de 3")
	}else if m%2==0{
		fmt.Println(m,"é multiplo somente de 2")
	}else if m%3==0{
		fmt.Println(m,"é multiplo somente de 3")
	}else{
		fmt.Println(m,"não é multiplo nem de 2 e nem de 3")
	}
	fmt.Println()
	for m= range 11{
		if m%2==0 && m%3==0{
			fmt.Println(m,"é multiplo de 2 e de 3")
		}else if m%2==0{
			fmt.Println(m,"é multiplo somente de 2")
		}else if m%3==0{
			fmt.Println(m,"é multiplo somente de 3")
		}else{
			fmt.Println(m,"não é multiplo nem de 2 e nem de 3")
		}
	}
}