package memory

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/jmjoy/render"
)

type memory struct {
	dir     string
	commons []string
	pool    map[string]*template.Template
}

func (this *memory) Init(config interface{}) error {
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

func (this *memory) AddCommonTpl(names ...string) error {
	for i := range names {
		names[i] = filepath.Join(this.dir, names[i])
	}
	this.commons = append(this.commons, names...)
	return nil
}

func (this *memory) GetTemplate(names ...string) (*template.Template, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("file: must give at least one file")
	}

	tplName := filepath.Base(names[0])
	for i := range names {
		names[i] = filepath.Join(this.dir, names[i])
	}
	names = append(names, this.commons...)
	seq := strings.Join(names, "%#@!!@#%")

	t, ok := this.pool[seq]
	if ok {
		return t, nil
	}

	t, err := template.New(tplName).ParseFiles(names...)
	if err != nil {
		return nil, err
	}
	this.pool[seq] = t
	return t, nil
}

func init() {
	driver := new(memory)
	driver.pool = make(map[string]*template.Template)
	render.Resigtry("memory", driver)
}
