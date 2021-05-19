package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	fmt.Println("Starting Backend")

	r := mux.NewRouter()
	r.HandleFunc("/branch", ReturnBranchNameHandler)
	http.Handle("/", r)
}

func ReturnBranchNameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "branch": "testing" }`)
}
