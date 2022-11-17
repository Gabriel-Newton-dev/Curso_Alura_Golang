package main

import "fmt"

type Nubank struct {
	Nome  string
	Conta int
	Saldo float64
}

func (c *Nubank) Depositar(valorDoDespoito float64) (string, float64, string) {
	if valorDoDespoito > 0 {
		c.Saldo += valorDoDespoito
		return "Deposito realizado com sucesso, seu saldo é de R$", c.Saldo, "Reais"
	} else {
		return "Não foi possível realizar deposito com valor negativo, seu saldo é de R$", c.Saldo, "Reais"
	}

}

func (c *Nubank) tranferencia(valorDaTransferencia float64, contaDestino *Nubank) (bool, string, float64) {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true, "Transferência realizada com sucesso o saldo atual é de R$", c.Saldo
	} else {
		return false, "Tranferência não realizada", valorDaTransferencia
	}
}

// func (d *DadosConta) tranferencia(valorDaTransferencia float64, contaDestino *DadosConta) bool {
// 	if valorDaTransferencia < d.Saldo {
// 		d.Saldo -= valorDaTransferencia
// 		contaDestino.Deposito((valorDaTransferencia))
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {

	contaDoAnderson := Nubank{"Anderson", 0333, 1000}
	contaDaGisele := Nubank{"Gisele", 0607, 300}

	fmt.Println(contaDoAnderson.Depositar(200))
	fmt.Println(contaDaGisele.tranferencia(100, &contaDoAnderson))
	fmt.Printf("Saldo atualizado da conta do %s após a transferência é de %v Reais\n", contaDoAnderson.Nome, contaDoAnderson.Saldo)

}
