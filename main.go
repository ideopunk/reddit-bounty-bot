package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func post(client *reddit.Client, title string, URL string) (string, error) {
	post, _, err := client.Post.SubmitLink(ctx,
		reddit.SubmitLinkRequest{Subreddit: "test", Title: title, URL: URL})

	if err != nil {
		return "", fmt.Errorf("could not submit post %v successfully: %w", title, err)
	}

	return post.URL, nil
}

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

	client, _ := reddit.NewReadonlyClient()
	posts, _, err := client.Subreddit.TopPosts(context.Background(), "test", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 5,
		}, Time: "all"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received %d posts.\n", len(posts))
	for i, post := range posts {
		println(i)
		fmt.Printf(post.Title)
	}
}
