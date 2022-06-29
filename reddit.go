package main

import (
	"context"
	"fmt"
	"os"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()

func getRedditPostingClient() (*reddit.Client, error) {
	ID := os.Getenv("REDDIT_APP_ID")
	secret := os.Getenv("REDDIT_APP_SECRET")
	username := os.Getenv("REDDIT_USERNAME")
	password := os.Getenv("REDDIT_PASSWORD")

	credentials := reddit.Credentials{ID: ID, Secret: secret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, fmt.Errorf("could not generate reddit client: %w", err)
	}

	return client, nil
}

func post(client *reddit.Client, posts []Post) ([]string, error) {
	urls := make([]string, len(posts))

	for _, post := range posts {
		fmt.Printf("%+v\n", post)

		// publishedPost, _, err := client.Post.SubmitLink(ctx,
		// 	reddit.SubmitLinkRequest{Subreddit: "test", Title: post.Title, URL: post.URL})

		// if err != nil {
		// 	return nil, fmt.Errorf("could not submit post %v successfully: %w. did however post successfully here: %v", post.Title, err, urls)
		// }

		// urls = append(urls, publishedPost.URL)
	}

	return urls, nil
}

// func testGet(client *reddit.Client) error {
// 	posts, _, err := client.Subreddit.TopPosts(context.Background(), "test", &reddit.ListPostOptions{
// 		ListOptions: reddit.ListOptions{
// 			Limit: 5,
// 		}, Time: "all"})

// 	if err != nil {
// 		return fmt.Errorf("could not get top posts: %w", err)
// 	}
// 	fmt.Printf("Received %d posts.\n", len(posts))

// 	return nil
// }
