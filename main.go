package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		BS          string `json:"bs"`
	} `json:"company"`
	Img string `json:"img"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Open the JSON file
	file, err := os.Open("user.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode the JSON file into a User slice
	var users []User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the User slice to JSON format
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers for CORS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow all origins, change as needed
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Allow only GET requests, change as needed

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUserHandler)
	http.Handle("/", r)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
