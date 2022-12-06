package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MilkShake struct {
	Sabor        string
	Valor        float64
	ComGranulado bool
}

func (m MilkShake) ColocarCalda() string {
	if m.ComGranulado == true {
		return "Adicionando Calda"
	} else {
		return "Não foi adicionado calda"
	}

}

func main() {

	Kibon := MilkShake{"Ovomaltine", 13.99, true}
	fmt.Println(Kibon.ColocarCalda())

	kibonJson, err := json.Marshal(Kibon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(kibonJson))
}
