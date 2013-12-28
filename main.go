package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "index.html", "template file")
	flag.Parse()
}

func main() {
	t, err := template.ParseFiles("base.html", file)
	if err != nil {
		fmt.Println("template parse error: ", err)
		return
	}
	err = t.Execute(os.Stdout, "")
	if err != nil {
		fmt.Println("template executing error: ", err)
		return
	}
}
