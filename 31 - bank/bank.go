package main

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		fmt.Printf("O valor do saque foi de R$ %v Reais, o Saldo atualizado é de R$ %v Reais\n", valorDoSaque, c.Saldo)
	} else {
		fmt.Println("Saldo Insuficiente")
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		fmt.Printf("O valor do Deposito foi de R$ %v Reais na conta de %s, seu Saldo atualizado é de R$ %v Reais\n", valorDoDeposito, c.Titular, c.Saldo)
	} else {
		fmt.Println("Não é possível depositar valor negativo")
	}
}

func (c *ContaCorrente) Transferencia(valorDaTransferencia float64, contaDestino *ContaCorrente) {
	podeTransferir := valorDaTransferencia < c.Saldo && valorDaTransferencia > 0
	if podeTransferir {
		c.Saldo -= valorDaTransferencia
		contaDestino.Saldo += valorDaTransferencia
		fmt.Printf("Valor de R$ %v Reais, foi transferido com sucesso, para conta do %v.\n", valorDaTransferencia, contaDestino.Titular)
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
	fmt.Println("R$", contaDaJurema.Saldo, "Reais")
	fmt.Println("R$", contaDoGuilherme.Saldo, "Reais")
}
