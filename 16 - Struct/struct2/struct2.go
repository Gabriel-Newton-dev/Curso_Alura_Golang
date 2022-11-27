package main

import "fmt"

type Carros struct {
	Marca  string
	Modelo string
	Valor  float64
	Turbo  bool
}

type Caminhonetes struct {
	Carros
	Importada           bool
	QuantidadeDeLugares int
}

func main() {

	Ram := Caminhonetes{
		Carros:              Carros{"Dodge Ram", "Ram 1500", 499.899, true},
		Importada:           true,
		QuantidadeDeLugares: 5,
	}

	Cruze := Carros{"Cruze", "Ltz", 100.799, true}

	fmt.Println(Ram, " | ", Cruze)

}
