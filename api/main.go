package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	a := App{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password='%s' "+
	// 	"dbname=%s sslmode=disable",
	// 	os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("LOGIN"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	// log.Info(psqlInfo)

	// a.Initialize(psqlInfo)
	a.Initialize()
	a.Run(":9001")

}
