package blogposts_test

import (
	"errors"
	"github.com/quzhi1/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

type Testcase struct {
	filename string
	markdown string
	post     blogposts.Post
}

func TestNewBlogPosts(t *testing.T) {
	serializedPost1 := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}
	serializedPost2 := blogposts.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"rust", "borrow-checker"},
		Body: `B
L
M`,
	}
	testcase1 := Testcase{
		filename: "hello world.md",
		markdown: `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`,
		post: serializedPost1,
	}
	testcase2 := Testcase{
		filename: "hello-world2.md",
		markdown: `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`,
		post: serializedPost2,
	}

	cases := []Testcase{
		testcase1,
		testcase2,
	}

	mockFs := fstest.MapFS{}

	for _, testcase := range cases {
		mockFs[testcase.filename] = &fstest.MapFile{Data: []byte(testcase.markdown)}
	}

	posts, err := blogposts.NewPostsFromFS(mockFs)
	if err != nil {
		t.Fatal(err)
	}

	for i, testcase := range cases {
		assertPost(t, posts[i], testcase.post)
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
