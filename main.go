package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

    log.Println("starting server")

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	fs := http.FileServer(http.Dir("."))
	http.Handle("/style.css", fs)

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl.Execute(w, nil)
	})

    addr := fmt.Sprintf("%s:%s", "127.0.0.1", "8080")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
