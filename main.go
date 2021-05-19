package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {

}

func main() {
	fmt.Println("Starting Backend")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			os.Getenv("REACT_APP_BACKEND"),
			os.Getenv("REACT_APP_FRONTEND"),
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"id-token", "content-type"},
	})

	r := mux.NewRouter()
	r.HandleFunc("/branch", ReturnBranchNameHandler)

	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

func ReturnBranchNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "branch": "main" }`)
}
