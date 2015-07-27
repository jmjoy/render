package main

import (
	"html/template"
	"os"
)

func main() {
	tpl, err := template.New("index.html").ParseFiles("index.html", "base.html")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"Content": "What the fuck!",
		"Title":   "Title Me",
	}
	tpl.Execute(os.Stdout, data)
}
