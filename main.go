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
	// Locations    string   `json:"locations"`
	// Concertdates string   `json:"concertDates"`
	// Relations    string   `json:"relations"`
}
type Relation struct {
	Index []RelationEntry `json:"index"`
}

type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type ArtistDetailPage struct {
	Artist   Artist
	Relation RelationEntry
}

func artistDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/artist/"):]
	artists, artisterr := getArtistDetails()
	datesLocations, datesLocationserr := getArtistRelations()
	if (artisterr != nil) || (datesLocationserr != nil) {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	var selected Artist
	var selectedRelation RelationEntry
	for _, artist := range artists {
		if fmt.Sprintf("%d", artist.ID) == idStr {
			selected = artist
			break
		}
	}
	for _, artistdateLocation := range datesLocations.Index {
		if fmt.Sprintf("%d", artistdateLocation.ID) == idStr {
			selectedRelation = artistdateLocation
			break
		}
	}
	if selected.ID == 0 {
		http.NotFound(w, r)
		return
	}
	tmpl, tmplerr := template.ParseFiles("templates/artist_detail.html")
	if tmplerr != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	data := ArtistDetailPage{
		Artist:   selected,
		Relation: selectedRelation,
	}

	tmpl.Execute(w, data)
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
func getArtistRelations() (Relation, error) {
	rel, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return Relation{}, err
	}
	defer rel.Body.Close()
	var relations Relation
	err = json.NewDecoder(rel.Body).Decode(&relations)
	if err != nil {
		return Relation{}, err
	}
	return relations, nil
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
