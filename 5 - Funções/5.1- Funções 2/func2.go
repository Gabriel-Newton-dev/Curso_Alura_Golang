package main

import "fmt"

func main() {

	a := 1
	b := 6

	fmt.Println("Os valores A e B multiplicado é : ", multiplicar(a, b))
	fmt.Println("Os valores A e B somados é : ", somar(a, b))
	quemEmaior(a, b)
	numeroDivisivelPorCinco()

}

func somar(n1, n2 int) int {

	return n1 + n2
}

func multiplicar(n1, n2 int) int {
	return n1 * n2
}

func quemEmaior(n1, n2 int) {

	var mensagem string

	if n1 < n2 {
		mensagem = "A número da letra A é menor do que da letra B"
	} else if n1 == n2 {
		mensagem = "A número da letra A é igual da letra B"
	} else {
		mensagem = "A número da letra é maior do que da letra B"
	}
	fmt.Println(mensagem)

}

func numeroDivisivelPorCinco() {
	for i := 0; i <= 50; i++ {
		if i%5 == 0 {
			fmt.Println(i, "é um número divisivel por 5")
		}
	}
}
