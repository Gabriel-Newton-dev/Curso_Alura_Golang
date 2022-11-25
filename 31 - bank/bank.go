package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		fmt.Printf("O valor do saque foi de R$ %v Reais, o saldo atualizado é de R$ %v Reais\n", valorDoSaque, c.saldo)
	} else {
		fmt.Println("Saldo Insuficiente")
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		fmt.Printf("O valor do Deposito foi de R$ %v Reais, seu saldo atualizado é de R$ %v Reais\n", valorDoDeposito, c.saldo)
	} else {
		fmt.Println("Não é possível depositar valor negativo")
	}
}

func (c *ContaCorrente) Transferencia(valorDaTransferencia float64, contaDestino *ContaCorrente) {
	podeTransferir := valorDaTransferencia < c.saldo && valorDaTransferencia > 0
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		fmt.Printf("Valor de R$ %v Reais, foi transferido com sucesso, para conta do %v.\n", valorDaTransferencia, contaDestino.titular)
	} else {
		fmt.Println("Transferência não realizada")
	}
}

func main() {

	contaDoGuilherme := ContaCorrente{"Guilherme Dias", 1612, 256969, 187.60}

	contaDaJurema := ContaCorrente{"Jurema", 222, 111222, 500}

	contaDoGuilherme.Sacar(1000)
	contaDaJurema.Sacar(1000)
	contaDoGuilherme.Depositar(13)
	contaDaJurema.Transferencia(200, &contaDoGuilherme)
	fmt.Println("R$", contaDaJurema.saldo, "Reais")
	fmt.Println("R$", contaDoGuilherme.saldo, "Reais")
}
