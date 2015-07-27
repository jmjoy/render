package render_test

import (
	"html/template"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	tpl, err := template.New("index.html").ParseFiles("./_testdata/index.html", "./_testdata/base.html")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"Content": "What the fuck!",
		"Title":   "Title Me",
	}
	tpl.Execute(os.Stdout, data)
}
