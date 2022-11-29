package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cadastro struct {
	Nome             string `json:"Nome"`
	Idade            int    `json:"Idade"`
	CPF              string `json:"Cpf"`
	NumeroDaCarteira int    `json:"Numero da carteira"`
}

type Endereco struct {
	Cadastro
	Rua    string `json:"Rua"`
	Numero int    `json:"Numero"`
	Bairro string `json:"Bairro"`
	Estado string `json:"Estado"`
}

func main() {

	CadastroDoRoberto := Endereco{
		Cadastro: Cadastro{"Roberto Silva", 58, "123.456.789-12", 56437},
		Rua:      "Rua Nova Holanda",
		Numero:   10,
		Bairro:   "Maré",
		Estado:   "Rio de Janeiro",
	}

	CadastroDoMiudilho := Endereco{
		Cadastro: Cadastro{"Felipe Felix", 37, "456-789-123-90", 17122},
		Rua:      "Favela da Vila do João",
		Numero:   171,
		Bairro:   "Bonsucesso",
		Estado:   "Rio de Janeiro",
	}

	robertoJson, err := json.Marshal(CadastroDoRoberto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(robertoJson))

	miudoJson, err := json.Marshal(CadastroDoMiudilho)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(miudoJson))
}
