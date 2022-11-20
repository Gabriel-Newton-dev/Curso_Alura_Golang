package main

import "fmt"

func main() {

	var nome = "Gabriel"

	switch nome {
	case "João":
		fmt.Println("Meu nome é o João")
	case "Gabriel":
		fmt.Println("Meu nome é o Gabriel")
	case "Daniel":
		fmt.Println("Meu nome é o Daniel")
	default:
		fmt.Println("entrada Inválida.")
	}

}
