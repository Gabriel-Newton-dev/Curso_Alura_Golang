package main

import (
	"fmt"
	"os"
)

var entrarSairPrograma string

func main() {

	fmt.Println("Por favor digite a opção desejada: \n1- Para entrar no programa.\n0 - Para sair do programa. ")
	fmt.Scan(&entrarSairPrograma)

	switch entrarSairPrograma {
	case "1":
		fmt.Println("Entrando no programa...")
	case "0":
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida, saindo do programa")
		os.Exit(-1)
	}

}
