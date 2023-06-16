package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lamsalsangam/Go/tree/master/rssaggregator/vendor/github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	fmt.Println("Port:", portString)
}
