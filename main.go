package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	PORT := os.Getenv("PORT")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(PORT, router))
}
