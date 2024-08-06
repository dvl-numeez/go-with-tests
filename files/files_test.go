package files

import (
	"testing"
	"testing/fstest"

	"github.com/dvl-numeez/go-with-tests/files/blogposts"
)

func TestNewBlogPost(t *testing.T) {
	t.Run("test the new blog function",func(t *testing.T){
		fs:=fstest.MapFS{
			"hello_world_1.md":{Data: []byte("Title: Post 1")},
			"hello_world_2.md":{Data: []byte("Title: Post 2")},
		}
		posts,err:=blogposts.NewPostsFromFs(fs)
		if err!=nil{
			t.Fatal(err)
		}
		got:=posts[0]
		want:=blogposts.Post{Title:"Post 1"}
		if got!=want{
			t.Errorf("got %+v, want %+v", got, want)
		}
		if len(posts)!=len(fs){
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
	
	t.Run("Check the failing of function",func(t *testing.T){
		_,err:=blogposts.NewPostsFromFs(blogposts.StubFileSystem{})
		if err==nil{
			t.Error("It should have failed")
		}
	})
}