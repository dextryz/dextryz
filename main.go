package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    addr := fmt.Sprintf("%s:%s", "0.0.0.0", port)

    log.Printf("serving on addr: %s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
