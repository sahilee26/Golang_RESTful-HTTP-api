
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User is a struct that represents a User in pur application
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// Post is struct that groups all necesaary fields into a single unit
type Post struct {
	// using json tags
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/add", addItem).Methods("POST")

	http.ListenAndServe(":5000", router)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	//get Item value from JSON Body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
