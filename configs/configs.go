package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Println(err)
	}
}