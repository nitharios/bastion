package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	router "github.com/nitharios/bastion/api"
)

func main() {
	PORT := os.Getenv("PORT")

	router := router.New()

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
