package main

import "fmt"

func main() {
	var nome string = "Vini"
	//outra maneira de usar variaveis (declaraçao implicita/inferencia de tipos)
	nome2 := "Isa"
	fmt.Println(nome, nome2)

	var (
		nome3 string = "Nicolas"
		nome4 string = "Samir"
	)

	fmt.Println(nome3, nome4)

	nome5, nome6 := "Nilda", "Sandra"

	fmt.Println(nome5, nome6)

}
