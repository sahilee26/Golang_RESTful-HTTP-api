package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	//  sending data over route with a dynamic variable with identifier data
	router.HandleFunc("/add/{item}", addItem)

	http.ListenAndServe(":5000", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"555"})
}

func addItem(w http.ResponseWriter, r *http.Request) {
	routeVariable := mux.Vars(r)["item"]
	data = append(data, routeVariable)

	json.NewEncoder(w).Encode(data)
}
