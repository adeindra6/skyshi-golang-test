package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	localServer := "http://localhost:3030"
	fmt.Println(fmt.Sprintf("Server running on: %s", localServer))
	log.Fatal(http.ListenAndServe("localhost:3030", r))
}
