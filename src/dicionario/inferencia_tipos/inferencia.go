package main

import "fmt"

func main() {
	/*
	--- Variáveis com declaração de tipos

	var nome string = "Glauco"
	var idade int = 46
	var versao float32 = 3.2
	*/

		/*
	--- Variáveis com inferencia de tipos
	
	var nome  = "Glauco"
	var idade  = 46
	var versao  = 3.2
	*/

	// declaração da variável com atribuiição e inferencia de tipo
	nome:="Glauco"
	idade:="46"
	versao:=3.2

	fmt.Println(`Meu nome é`, nome, "e minha idade é", idade)
	fmt.Println(`Estou usando o program de versão:`, versao)
}