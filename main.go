package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	fmt.Println("Starting Backend")

	r := mux.NewRouter()
	r.HandleFunc("/branch", ReturnBranchNameHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func ReturnBranchNameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "branch": "main" }`)
}
