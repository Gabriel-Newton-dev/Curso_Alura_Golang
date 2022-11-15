package main

import "fmt"

type ContaCorrrente struct {
	Nome  string
	Conta int
	Saldo float64
}

func main() {

	ContaDaCristiane := ContaCorrrente{"Cristiane Macedo", 256969, 90.786}
	fmt.Println(ContaDaCristiane)

	ContaDoHulk := ContaCorrrente{"Hulk", 221712, 123.89}
	fmt.Println(ContaDoHulk)

	// criei um ponteiro para

	var ContaDaCris *ContaCorrrente
	ContaDaCris = new(ContaCorrrente) // temos que usar o comando new para além do ponteiro, referenciar que é a criação de uma nova Conta Corrente.
	ContaDaCris.Nome = "Cristina"
	ContaDaCris.Conta = 256969
	ContaDaCris.Saldo = 40.678

	fmt.Println(*ContaDaCris)
}
