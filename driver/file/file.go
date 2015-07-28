package driver

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/jmjoy/render"
)

type File struct {
	dir     string
	commons []string
}

func (this *File) Init(config interface{}) error {
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

func (this *File) AddCommonTpl(names ...string) error {
	for i := range names {
		names[i] = filepath.Join(this.dir, names[i])
	}
	this.commons = append(this.commons, names...)
	return nil
}

func (this *File) GetTemplate(names ...string) (*template.Template, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("file: must give at least one file")
	}

	tplName := filepath.Base(names[0])
	names = append(names, this.commons...)
	return template.New(tplName).ParseFiles(names...)
}

func init() {
	render.Resigtry("file", new(File))
}
