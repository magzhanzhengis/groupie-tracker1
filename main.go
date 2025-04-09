package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Joke struct to hold the joke data
type Artist struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Members []string `json:"members"`
}

// Fetch random joke from API
func getArtistDetails() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

// Home handler to serve the page åand display the joke
func homepageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	// Passing the joke to the template
	tmpl.Execute(w, nil)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := getArtistDetails()
	if err != nil {
		http.Error(w, "Failed to fetch artist", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	// Passing the joke to the template
	tmpl.Execute(w, artists)
}

func main() {
	// Route to handle the homepage
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/artists", artistsHandler)

	// Start server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
