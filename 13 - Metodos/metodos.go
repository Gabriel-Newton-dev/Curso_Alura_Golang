package main

import "fmt"

type ContaCorrente struct {
	Nome  string
	Saldo float64
}

// Nessa funcao temos um metodo ( c Contacorrente) temos uma funcao depositar, que recebe como parametro deposito que é
// float64 e irá ter como retorno 1 string, 1 float64 e  outra string,
func (c ContaCorrente) Depositar(deposito float64) (string, float64, string) {
	if deposito > 0 {
		c.Saldo += deposito
		return "Deposito realizado com sucesso, valor atual do saldo R$", c.Saldo, "Reais"
	} else {
		return "Não é possível realizar o deposito com valor negativo, seu saldo atual é R$", c.Saldo, "Reais"
	}

}

func (c *ContaCorrente) Sacar(Saque float64) (string, float64, string) {
	if Saque > 0 && Saque < c.Saldo {
		c.Saldo -= Saque
		return "Deposito realizado com sucesso, seu saldo é de R$", c.Saldo, "Reais"
	} else {
		return "Não foi possível realizar o deposito, seu saldo atual é de R$", c.Saldo, "Reais"
	}

}

func main() {

	contaDoGivanildo := ContaCorrente{"Givanildo", 500}
	fmt.Println(contaDoGivanildo.Depositar(500))

	fmt.Println(contaDoGivanildo.Sacar(100))

}
