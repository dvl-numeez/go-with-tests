package templating

import (
	"bytes"
	"testing"
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
	t.Run("converts a post to html",func(t *testing.T){
		buff:=&bytes.Buffer{}
		err:=Render(buff,aPost)
		if err!=nil{
			t.Fatal(err)
		}
		got:=buff.String()
		want:=`<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}