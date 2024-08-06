package blogposts

import (
	"errors"
	"io"
	"io/fs"
)

type StubFileSystem struct{

}
func (sf StubFileSystem)Open(name string) (fs.File, error){
	return nil,errors.New("this is a custom test to check the failing")
}


type Post struct{
	Title string
}

func NewPostsFromFs(fileSystem fs.FS)([]Post,error){
	dir,err:=fs.ReadDir(fileSystem,".")
	if err!=nil{
		return nil,err
	}
	var posts []Post
	for _,f:=range dir{
		post,err:=getPost(fileSystem,f)
		if err!=nil{
			return nil,err
		}
		posts = append(posts, post)
	}
	return posts,nil
}

func getPost(fileSystem fs.FS,file fs.DirEntry)(Post,error){
		fs,err:=fileSystem.Open(file.Name())
		if err!=nil{
			return Post{}, err
		}
		defer fs.Close()
		postData,err:=io.ReadAll(fs)
		if err!=nil{
			return Post{}, err
		}
		return Post{Title: string(postData)[7:]},nil
}