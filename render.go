package render

import (
	"fmt"
	"html/template"
	"sort"
)

type Driver interface {
	Init(config interface{}) error
	AddCommonTpl(names ...string) error
	GetTemplate(names ...string) (*template.Template, error)
	Delims(left, right string)
}

func Resigtry(name string, d Driver) {
	if d == nil {
		panic("render: Register deiver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("render: Register called twice for driver " + name)
	}
	drivers[name] = d
}

func Drivers() []string {
	var list []string
	for name := range drivers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func New(driverName string, config interface{}) (Driver, error) {
	d, ok := drivers[driverName]
	if !ok {
		return nil, fmt.Errorf("render: unknown driver %q (forgotten import?)", driverName)
	}

	err := d.Init(config)
	if err != nil {
		return nil, err
	}

	return d, nil
}

var drivers = make(map[string]Driver)
