package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice := []string{"Gabriel", "Douglas", "Diego", "Roberta"}
	fmt.Println(slice)
	fmt.Println("A capacidade do meu slice é de :", cap(slice))
	slice = append(slice, "Bernadete")
	fmt.Println(slice)

	// para saber qual o tipo utilizamos o reflect.Typeof(nome do arquivos que desejamos saber)
	fmt.Println(reflect.TypeOf(slice))
	// quando vc coloca o append ele dobra o valor, pq na verdade o slice é uma abstracao do array.
	fmt.Println("A capacidade do meu slice é de :", cap(slice))
	fmt.Println("O tamanho do meu slice é", len(slice), "itens")

}
