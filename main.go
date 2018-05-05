package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func loadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadDotenv()

	se, err := mgo.Dial(os.Getenv("MONGO_HOST") + ":27017")
	if err != nil {
		panic(err)
	}
	defer se.Close()
	se.SetMode(mgo.Monotonic, true)

	r := gin.Default()

	SetItemRouter(r, se)

	r.Run()
}
