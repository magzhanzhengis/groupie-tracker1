package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Joke struct to hold the joke data
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	Firstalbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Concertdates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func artistDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/artist/"):]
	artists, err := getArtistDetails()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	var selected Artist
	for _, artist := range artists {
		if fmt.Sprintf("%d", artist.ID) == idStr {
			selected = artist
			break
		}

	}
	if selected.ID == 0 {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("templates/artist_detail.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, selected)
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
	http.HandleFunc("/artist/", artistDetailHandler)

	// Start server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
