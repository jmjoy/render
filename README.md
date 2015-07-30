# render

对Go的模板引擎的简单封装

## Features

 * use driver pattern like `database/sql`, for example, use `file` driver when develop or use `memory` when production
 * Too simple

## Tutorial

 ```Go
import (
	"github.com/jmjoy/render"
	_ "github.com/jmjoy/render/driver/file"
	// _ "github.com/jmjoy/render/driver/memory"
)

func TestRender(t *testing.T) {
		d, err := render.New("file", map[string]string{
			"dir": "path/to/views",
		})
		if err != nil {
			t.Fatal(err)
		}

		d.AddCommonTpl("base.html")
		tpl := template.Must(d.GetTemplate("index.html"))

        // use tpl as html/*template.Template
        // ...
	}
}
 ```

## version

v0.0.1

## License

MIT
