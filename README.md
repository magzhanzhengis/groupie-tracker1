

# 🎸 Groupie Tracker

Groupie Tracker is a web application that displays information about musical artists and their tour dates. It fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/) and presents it in an intuitive interface. Users can browse artists, view their details, and explore tour date locations across the world.

## 🌐 Features

- 📜 **Artist List View**: Browse through a list of musical artists with basic info.
- 🧑‍🎤 **Artist Detail View**: View detailed information including:
  - Name, image, and members
  - First album release date
  - Creation date
  - Concert locations and corresponding dates


## 🏗️ Project Structure

```
groupie-tracker/
│
├── main.go                 # Entry point of the Go web server
├── templates/
│   ├── homepage.html       # Main landing page template
│   └── artist_detail.html  # Detailed artist info page
├── static/
│   └── style.css           # Optional CSS for styling
└── README.md               # You're reading it!
```

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (v1.18 or later)

### Installation & Running

```bash
git clone https://github.com/yourusername/groupie-tracker.git
cd groupie-tracker
go run .
```

Then open your browser and visit: [http://localhost:8080](http://localhost:8080)

## 🔗 API Reference

Data is fetched from the official Groupie Tracker API:
- **Artists**: `https://groupietrackers.herokuapp.com/api/artists`
- **Relations**: `https://groupietrackers.herokuapp.com/api/relation`
- **Locations**: `https://groupietrackers.herokuapp.com/api/locations`
- **Dates**: `https://groupietrackers.herokuapp.com/api/dates`

## 🛠️ Technologies Used

- **Go** – for the backend web server
- **HTML Templates** – to render dynamic views
- **net/http** – Go standard library for web development

## ✅ TODO / Improvements

- [ ] Add search functionality
- [ ] Implement location map using Leaflet.js or Google Maps API
- [ ] Add filters for year, location, or band members
- [ ] Mobile-friendly styling with responsive CSS

## 👨‍💻 Author

Made by **Magzhan Zhenis** as a portfolio/web development project.
