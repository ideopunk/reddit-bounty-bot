package main

import (
	"context"
	"fmt"
	"log"

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

func testGet(client *reddit.Client) {
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
