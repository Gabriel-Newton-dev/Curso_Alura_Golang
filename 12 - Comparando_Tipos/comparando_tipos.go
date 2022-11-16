package main

import "fmt"

type ContaBanco struct {
	Nome        string
	Agencia     int
	NumeroConta int
	Saldo       float64
}

func main() {

	ContaGuilherme := ContaBanco{"Guilherme", 587, 445334, 8.560}

	ContaGuilherme2 := ContaBanco{"Guilherme", 587, 445334, 8.560}

	fmt.Println(ContaGuilherme, ContaGuilherme2)

	// estamos comparando o conteudo
	fmt.Println("iremos comprar se a conta 1 do Guilherme é igual a conta 2, sendo assim o comparativo é:", ContaGuilherme == ContaGuilherme2)

	var ContaDaCris *ContaBanco
	ContaDaCris = new(ContaBanco) // temos que usar o comando new para além do ponteiro, referenciar que é a criação de uma nova Conta Corrente.
	ContaDaCris.Nome = "Cristina"
	ContaDaCris.NumeroConta = 256969
	ContaDaCris.Saldo = 40.678

	var ContaDaCris2 *ContaBanco
	ContaDaCris = new(ContaBanco) // temos que usar o comando new para além do ponteiro, referenciar que é a criação de uma nova Conta Corrente.
	ContaDaCris.Nome = "Cristina"
	ContaDaCris.NumeroConta = 256969
	ContaDaCris.Saldo = 40.678

	// Por mais que sejam identicos, como se trata de um ponteiro, o numero de
	fmt.Println(ContaDaCris == ContaDaCris2)

	// agora quando comparamos o conteudo do ponteiro

	fmt.Println(*ContaDaCris == *ContaDaCris2)

}
