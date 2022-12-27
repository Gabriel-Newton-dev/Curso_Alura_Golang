package main

import (
	"fmt"
	"net/http"
)

type Pessoas struct {
	Nome                    string
	Idade                   int
	NacionalidadeBrasileira bool
}

func main() {

	Sofia := Pessoas{"Sofia", 25, true}
	Rita := Pessoas{"Rita", 30, false}

	fmt.Println(Sofia, Rita)

}

func RetornaTodasAsPessoas(w http.ResponseWriter, r *http.Request) {

}
