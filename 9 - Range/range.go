package main

import "fmt"

func main() {

	carros := []string{"Cruze", "Civic", "Corolla", "Hillux", "Siena", "Vectra"}

	for key, value := range carros {
		fmt.Println(key, value)
	}
	fmt.Println()
	number := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for key, value := range number {
		fmt.Println(key, value)
	}

}
