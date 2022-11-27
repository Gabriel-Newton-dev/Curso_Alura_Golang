package main

import (
	"encoding/json"
	"fmt"
)

type Console struct {
	Marca       string  `json:"Marca"`
	Modelo      string  `json:"Modelo"`
	Valor       float64 `json:"Valor"`
	NovaGeração bool    `json:"Nova Geração"`
}

func main() {

	WiiU := Console{"Wii U", "Nintendo", 1200, false}

	Ps5 := Console{"PS5", "Sony", 5789.90, true}

	WiiUJson, err := json.Marshal(WiiU)
	if err != nil {
		fmt.Println("Erro Json WiiU", err)
	}
	fmt.Println(string(WiiUJson))

	Ps5Json, err := json.Marshal(Ps5)
	if err != nil {
		fmt.Println("Erro Json Ps5", err)
	}
	fmt.Println(string(Ps5Json))

}
