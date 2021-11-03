package main

import (
	"github.com/quzhi1/blogposts"
	"log"
	"os"
)

func main() {
	dirFs := os.DirFS("posts")
	posts, err := blogposts.NewPostsFromFS(dirFs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", posts)
}
