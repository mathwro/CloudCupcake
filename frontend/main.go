package main

import (
	"html/template"
	"net/http"
	"log"
)

func main() {
	// Define the template with the HTML and CSS embedded
	tmpl := template.Must(template.ParseFiles("./layout/index.html"))

	// Serve the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Serve the CSS file separately
	http.HandleFunc("./layout/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// Start the server
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
