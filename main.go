package main

import (
	"fmt"
	"html/template"
	"os"
)

type Mapper struct {
	Name string
}

const tmplFile = "mapper.tmpl"

func main() {
	mapper := Mapper{
		Name: "What",
	}

	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%sMapper.ts", mapper.Name))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, mapper)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
