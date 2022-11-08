package main

import "fmt"

func main() {
	// exibeValores()
	// leComando()

	nome, idade := devolveNomeEIdade()
	fmt.Printf("Meu nome é %s e tenho %d anos de idade.", nome, idade)

	number, texto, VouF, _ := verificadorSites()
	fmt.Println(number, texto, VouF)

}

// funcao simples que não retorna parametros em GO, não precisa colocar nada, só o que deseja no bloco do código
func exibeValores() {
	nome := "Função simples em GO"
	fmt.Println(nome)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func devolveNomeEIdade() (string, int) {
	nome := "Gabriel"
	idade := 36
	return nome, idade
}

func verificadorSites() (int, string, bool, string) {
	return 10, "Verificador", true, "RJ"
}
