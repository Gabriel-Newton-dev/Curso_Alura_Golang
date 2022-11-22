package main

import "fmt"

func main() {

	animes := [3]string{"Dragon Ball", "Sailor Moon", "Yuyu Hakusho"}

	animesDoisPrimeiros := animes[:2]

	ultimoAnime := animes[2]

	fmt.Println(len(animes))
	fmt.Println(animesDoisPrimeiros)
	fmt.Println(ultimoAnime)

	// informe a capaciade e o tamanho da varaivel animesDoisPrimeiros usando printf

	fmt.Printf("len: %d, Cap: %d\n", len(animesDoisPrimeiros), cap(animesDoisPrimeiros))

}
