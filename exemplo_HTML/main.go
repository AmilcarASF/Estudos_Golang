package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("*.html"))

	usuario1 := usuario{"Amilcar", "emailteste@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", usuario1)
}

func main() {
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
