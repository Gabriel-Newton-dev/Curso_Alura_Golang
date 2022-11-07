package main

import "fmt"

func main() {
	exibeValores()
	leComando()

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
