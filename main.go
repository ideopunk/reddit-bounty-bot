package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	// GET
	campaigns, err := getCampaigns()

	if err != nil {
		log.Fatal("could not get campaigns: %w", err)
	}

	println(len(campaigns))
	// POST
	client, err := getRedditPostingClient()

	if err != nil {
		log.Fatal("could not generate posting client: %w", err)
	}

	url, err := post(client, campaigns)

	if err != nil {
		log.Fatal("could not get posts: %w", err)
	}

	println("success")
	fmt.Printf("%v", url)
}
