package render_test

import (
	"html/template"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.Must(
		template.New("index.html").ParseGlob("_testdata/*.html"), //ParseFiles("./_testdata/index.html", "./_testdata/base.html")
	)

	data := map[string]string{
		"Content": "What the fuck!",
		"Title":   "Title Me",
	}
	tpl.Execute(os.Stdout, data)
}
