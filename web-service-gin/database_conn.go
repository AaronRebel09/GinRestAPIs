package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Album represents a record in the albums table
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func main_album() {
	var err error

	// Connect to SQLite database
	db, err = sql.Open("sqlite3", "music.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	// Set up HTTP routes
	http.HandleFunc("/albums", getAlbumsHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getAlbumsHandler handles the GET request to retrieve albums
func getAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	albums, err := getAlbum()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve albums: %v", err), http.StatusInternalServerError)
		return
	}

	// Encode albums to JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func conversion() {
	strconv.Atoi()
	strconv.Itoa()
}

// getAlbums retrieves all albums from the database
func getAlbum() ([]Album, error) {
	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	// Check for errors from iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}