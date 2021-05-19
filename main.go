package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	fmt.Println("Hello world")

	r := mux.NewRouter()
	http.Handle("/", r)
}
