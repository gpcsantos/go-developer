package main

import "fmt"

func main(){
	// resto %
	for i := 1; i <= 10; i++ {
		// if i == 1{
		// 	fmt.Println("um")
		// }else if i==2{
		// 	fmt.Println("dois")
		// }else if i==3{
		// 	fmt.Println("três")
		// }else if i==4{
		// 	fmt.Println("quatro")
		// }else if i==5{
		// 	fmt.Println("cinco")
		// }else if i==6{
		// 	fmt.Println("seis")
		// }else if i==7{
		// 	fmt.Println("sete")
		// }else if i==8{
		// 	fmt.Println("oito")
		// }else if i==9{
		// 	fmt.Println("nove")
		// }else if i==10{
		// 	fmt.Println("dez")
		// }
		// fmt.Println()

		switch i{
		case 1: 
			fmt.Println("um")
		case 2:
			fmt.Println("dois")
		case 3: fmt.Println("três")
		case 4: fmt.Println("quatro")
		case 5: fmt.Println("cinco")
		case 6: fmt.Println("seis")
		case 7: fmt.Println("sete")
		case 8: fmt.Println("oito")
		case 9: fmt.Println("nove")
		case 10: fmt.Println("dez")
		}

	}
}