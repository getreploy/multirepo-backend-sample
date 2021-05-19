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
			os.Getenv("REPLOY_BACKEND"),
			os.Getenv("REPLOY_FRONTEND"),
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"id-token", "content-type"},
	})

	r := mux.NewRouter()
	r.HandleFunc("/health", OKHandler)
	r.HandleFunc("/branch", ReturnBranchNameHandler)

	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

func ReturnBranchNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "branch": "testing-1234" }`)
}

func OKHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
