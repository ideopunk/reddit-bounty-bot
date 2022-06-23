package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func getPostingClient() (*reddit.Client, error) {
	ID := os.Getenv("REDDIT_APP_ID")
	secret := os.Getenv("REDDIT_APP_SECRET")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	credentials := reddit.Credentials{ID: ID, Secret: secret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, fmt.Errorf("could not generate reddit client: %w", err)
	}
	return client, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// client, err := getPostingClient()
	// client, err := reddit.NewReadonlyClient()
	if err != nil {
		log.Fatal("could not generate client: %w", err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {

		// title, url, legit := "", "", false

		legit := true
		println(legit)
		// if legit {
		// 	// post(client, title, url)
		// 	testGet(client)
		// }

	})
	http.ListenAndServe(":3000", nil)
}
