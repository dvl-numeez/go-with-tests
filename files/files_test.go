package files

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/dvl-numeez/go-with-tests/files/blogposts"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
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

	t.Run("test the new blog function",func(t *testing.T){
		fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
		posts,err:=blogposts.NewPostsFromFs(fs)
		if err!=nil{
			t.Fatal(err)
		}
			assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})

	})
	
	t.Run("Check the failing of function",func(t *testing.T){
		_,err:=blogposts.NewPostsFromFs(blogposts.StubFileSystem{})
		if err==nil{
			t.Error("It should have failed")
		}
	})
}
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}