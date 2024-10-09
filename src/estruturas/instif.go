package main

import "fmt"

func main(){
	// resto %
	for i := 1; i <= 10; i++ {
		if i%2 == 0{
			fmt.Println(i, "é Par")
		}else{
			fmt.Println(i, "é Impar")
		}
	}
}