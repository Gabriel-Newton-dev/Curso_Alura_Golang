package main

import (
	"fmt"
	"net/http"
)

func main() {

	// carros := []string{"Cruze", "Civic", "Corolla", "Hillux", "Siena", "Vectra"}

	// for key, value := range carros {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println()
	// number := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// for key, value := range number {
	// 	fmt.Println(key, value)
	// }

	sites2 := []string{"https://www.globo.com", "https://www.google.com.br", "https://www.uol.com.br"}

	for _, value := range sites2 {
		// fmt.Println(value)
		TestaSites(value)
	}

}

func TestaSites(sites string) {
	resp, err := http.Get(sites)
	if err != nil {
		fmt.Println("O site apresentou problema")
	}
	if resp.StatusCode == 200 {
		fmt.Println("O site", sites, "foi carregado com sucesso.")
	} else {
		fmt.Printf("O site: %s, está com problema. Status code:%d\n", sites, resp.StatusCode)
	}

}
