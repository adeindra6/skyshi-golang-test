package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}

	r := mux.NewRouter()
	http.Handle("/", r)
	localServer := fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))
	fmt.Println(fmt.Sprintf("Server running on: %s", localServer))
	listenServer := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(listenServer, r))
}
