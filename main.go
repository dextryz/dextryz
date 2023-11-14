package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	fs := http.FileServer(http.Dir("."))
	http.Handle("/style.css", fs)

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
