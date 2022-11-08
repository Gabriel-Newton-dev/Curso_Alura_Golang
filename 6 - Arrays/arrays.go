package main

import "fmt"

var array [3]string

func main() {

	estados := devolveEstados()
	fmt.Println(estados)

	array[0] = "https://www.alura.com.br"
	array[1] = "Teste"
	array[2] = "Array em Golang"

	fmt.Println(array)
}

func devolveEstados() [4]string {
	var estados [4]string
	estados[0] = "RJ"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"

	return estados
}
