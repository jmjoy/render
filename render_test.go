package render_test

import (
	"html/template"
	"os"
	"testing"
)

func TestRender(t *testing.T) {

	tpl := template.Must(
		template.New("_testdata/index.html").ParseFiles("D:/gopath/src/github.com/jmjoy/render/_testdata/index.html", "D:/gopath/src/github.com/jmjoy/render/_testdata/base.html"),
	)

	data := map[string]string{
		"Content": "What the fuck!",
		"Title":   "Title Me",
	}
	tpl.Execute(os.Stdout, data)
}
