package main

import (
	"html/template"
	"net/http"
)

var AbrirPaginaHtml = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", LerPaginaHtml)
	http.ListenAndServe(":9000", nil)
}

func LerPaginaHtml(w http.ResponseWriter, r *http.Request) {

	AbrirPaginaHtml.ExecuteTemplate(w, "Index", nil)

}
