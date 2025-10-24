package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "example.com/hello/blogposts_test"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("error")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: Tag 1, Tag 2
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"Tag 1", "Tag 2"},
			Body: `Hello
World`,
		}

		assertPost(t, got, want)
	})

	t.Run("should show error on failing fs", func(t *testing.T) {
		fs := StubFailingFS{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Error("was expecting an error, didnt get one")
		}
	})
}

func assertPost(t testing.TB, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
