package templating

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)


func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}
	t.Run("converts a post to html",func(t *testing.T){
		buff:=&bytes.Buffer{}
		err:=postRenderer.Render(buff,aPost)
		if err!=nil{
			t.Fatal(err)
		}
		approvals.VerifyString(t,buff.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
	
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
	
	
}


func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}	
	)
	postRenderer,err:= NewPostRenderer()
	if err!=nil{
		b.Fatal(err)
	}
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		postRenderer.Render(io.Discard,aPost)
	}


	
}