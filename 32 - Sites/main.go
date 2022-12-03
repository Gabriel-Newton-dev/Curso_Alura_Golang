package main

import (
	"net/http"
	"text/template"
)

var chamarTemplate = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", call)
	http.ListenAndServe(":2000", nil)
}

func call(w http.ResponseWriter, r *http.Request) {
	chamarTemplate.ExecuteTemplate(w, "Index", nil)
}
