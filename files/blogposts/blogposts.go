package blogposts

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"

	"io"
	"io/fs"
)
const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "

)



type StubFileSystem struct{

}
func (sf StubFileSystem)Open(name string) (fs.File, error){
	return nil,errors.New("this is a custom test to check the failing")
}


type Post struct{
	Title string
	Description string
	Tags []string
	Body string
}

func NewPostsFromFs(fileSystem fs.FS)([]Post,error){
	dir,err:=fs.ReadDir(fileSystem,".")
	if err!=nil{
		return nil,err
	}
	var posts []Post
	for _,f:=range dir{
		post,err:=getPost(fileSystem,f.Name())
		if err!=nil{
			return nil,err
		}
		posts = append(posts, post)
	}
	return posts,nil
}

func getPost(fileSystem fs.FS,file string)(Post,error){
		fs,err:=fileSystem.Open(file)
		if err!=nil{
			return Post{}, err
		}
		defer fs.Close()
		return newPost(fs)
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	title:=readMetaLine(titleSeparator)
	description:=readMetaLine(descriptionSeparator)
	tags:= strings.Split(readMetaLine(tagsSeparator), ", ")
	body := readBody(scanner)

	return Post{
		Title: title  ,
		Description: description,
		Tags:   tags    ,
		Body: body,
	}, nil
}


func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}