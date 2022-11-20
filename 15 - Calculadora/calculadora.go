package main

import "fmt"

var number1 float64

var number2 float64

var operadores string

func main() {

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Digite a operacao que deseja fazer: + | - |  * | / :  ")
	fmt.Scan(&operadores)

	switch operadores {
	case "+":
		fmt.Println(number1 + number2)
		break
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	default:
		fmt.Println("Entrada inválida, saindo da calculadora")
	}

}
