package render_test

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/jmjoy/render"
	_ "github.com/jmjoy/render/driver/file"
	_ "github.com/jmjoy/render/driver/memory"
)

func TestRender(t *testing.T) {
	driverNames := [...]string{"file", "memory"}

	for _, v := range driverNames {
		d, err := render.New(v, map[string]string{
			"dir": "./_testdata",
		})
		if err != nil {
			t.Fatal(err)
		}

		d.AddCommonTpl("base.html")
		tpl := template.Must(d.GetTemplate("index.html"))

		buf := &bytes.Buffer{}

		tpl.Execute(buf, map[string]string{
			"Title":   "标题",
			"Content": "内容",
		})

		if testdata1 != buf.String() {
			t.Fatalf("Content not equal %q", v)
		}

		_ = template.Must(d.GetTemplate("index1.html"))
	}

}

var testdata1 = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>标题</title>

<link rel="stylesheet" href="hello.css" type="text/css" media="screen" charset="utf-8">

</head>
<body>

<h1>Hello world!</h1>
<p>内容</p>
</body>

<script type="text/javascript" charset="UTF-8">
function sayHello() {
alert("Hello world");
}
</script>
</html>
`
