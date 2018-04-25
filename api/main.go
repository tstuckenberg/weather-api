package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

//Entry point to application
func main() {

	a := App{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a.Initialize()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		a.Run(":9000")
	} else {
		a.Run(port)
	}

}
