package main

import (
	"html/template"
	"io"
	"net/http"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	// Parse layout and main content
	tmpl, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the full template
	// tmpl.ExecuteTemplate(w, "index", "artists", nil)
	tmpl.Execute(w, nil)
}
func artistsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/artists", artistsHandler)
	http.HandleFunc("/api/artists", apiHandler)

	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
