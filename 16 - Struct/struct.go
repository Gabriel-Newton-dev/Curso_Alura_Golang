package main

import "fmt"

type Desktop struct {
	marca       string
	modelo      string
	processador string
	valor       float64
}

func main() {

	lenovo := Desktop{"Lenovo", "x20", "I9 10 geração", 4980.40}
	fmt.Println(lenovo)

}
