package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	c.saldo -= valorDoSaque
	fmt.Printf("O valor do saque foi de R$ %v Reais, o saldo atualizado é de R$ %v Reais\n", valorDoSaque, c.saldo)
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	c.saldo += valorDoDeposito
	fmt.Printf("O valor do Deposito foi de R$ %v Reais, seu saldo atualizado é de R$ %v Reais", valorDoDeposito, c.saldo)

}

func main() {

	contaDoGuilherme := ContaCorrente{"Guilherme Dias", 1612, 256969, 187.60}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	contaDoGuilherme.Sacar(10)
	fmt.Println(contaDoGuilherme)
	contaDaBruna.Sacar(10)
	contaDoGuilherme.Depositar(1000)
}
