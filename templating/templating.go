package templating

import (
	"embed"
	
	"html/template"
	"io"
	"strings"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	
)


var (
	//go:embed "templates/*"
	postTemplates embed.FS
)
type Post struct{
	Title, Description, Body string
	Tags                     []string

}
type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

type PostRenderer struct {
	templ *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (p Post)SanitisedTitle()string{
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	return  r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer,posts []Post)error{
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}