package util

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cant load .env file with godotenv\n", err)
	}
}