package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
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
