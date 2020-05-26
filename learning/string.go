
// learning to send string 
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	http.ListenAndServe(":5000", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test"))
}
