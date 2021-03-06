package file

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/jmjoy/render"
)

type file struct {
	dir         string
	commons     []string
	rawTemplate *template.Template
	leftDelim   string
	rightDelim  string
}

func (this *file) Init(config interface{}) error {
	m, ok := config.(map[string]string)
	if !ok {
		return fmt.Errorf("file: only support config format `map[string]string` now")
	}

	dir, ok := m["dir"]
	if !ok {
		return fmt.Errorf("file: must specify dir of config")
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	this.dir = absDir
	return nil
}

func (this *file) AddCommonTpl(names ...string) error {
	for i := range names {
		names[i] = filepath.Join(this.dir, names[i])
	}
	this.commons = append(this.commons, names...)
	return nil
}

func (this *file) GetTemplate(names ...string) (*template.Template, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("file: must give at least one file")
	}

	tplName := filepath.Base(names[0])
	for i := range names {
		names[i] = filepath.Join(this.dir, names[i])
	}
	names = append(names, this.commons...)

	t := template.New(tplName)
	if this.leftDelim != "" && this.rightDelim != "" {
		t = t.Delims(this.leftDelim, this.rightDelim)
	}
	return t.ParseFiles(names...)
}

func (this *file) Delims(left, right string) {
	this.leftDelim = left
	this.rightDelim = right
}

func init() {
	render.Resigtry("file", new(file))
}
