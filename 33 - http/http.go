package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", LerPaginaHtml)
	http.ListenAndServe(":9000", nil)
}

func LerPaginaHtml(w http.ResponseWriter, r *http.Request) {

}
