package main

import "fmt"

var idade int

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	sliceString := []string{"Gabriel", "Gustavo", "Ingrid", "Marcela", "Douglas"}

	for i := 0; i < len(sliceString); i++ {
		fmt.Println(sliceString[i])
	}

}
