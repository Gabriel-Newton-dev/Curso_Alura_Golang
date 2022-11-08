package main

import "fmt"

var prato string

func main() {

	fmt.Println("Digite o seu prato favorito?\nP - Pizza\nH - Hamburguer\nB - Bife com fritas\nO - Outros")
	fmt.Scan(&prato)

	switch prato {
	case "P":
		fmt.Println("Calabresa ou Napolitana?")
	case "H":
		fmt.Println("Com queijou ou ovo?")
	case "B":
		fmt.Println("Batata frita ou Noisete?")
	case "O":
		fmt.Println("Infelimente só temos esse cardápio no momento")
	default:
		fmt.Println("Opção inválida")
	}

}
