package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

// tpl is our template
var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>

    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}

        <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>
`

// NewHandler is used to handle web requests
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err) // not a good idea, print err instead
	}
}

// JsonStory reads/decodes a json object
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story is what we will be using in our code
type Story map[string]Chapter

// Chapter represents the struct that contains our json data
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option organizes the data for the options portion in our chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
