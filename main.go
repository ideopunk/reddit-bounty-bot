package main

import (
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

	println(campaigns)

	// POST
	client, err := getRedditPostingClient()

	// if err := post(client, campaign.Name + ": " + campaign.Blurb, url); err != nil {
	if err := testGet(client); err != nil {
		log.Fatal("could not get posts: %w", err)
	}

	if err != nil {
		log.Fatal("could not generate client: %w", err)
	}

	println("success")
}
