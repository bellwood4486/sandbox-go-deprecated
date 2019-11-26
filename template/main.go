package main

import (
	"log"
	"os"
	"text/template"
)

const templ1 = `{
  "id": "{{.ID}}",
  "name": "{{.Name}}"
}
`

const templ2 = `{
  "age": {{.Age}}
}`

const templ3 = `{
  "unknown": {{.Unknown}}
}`

var reqBody1 = template.Must(template.New("request body 1").
	Parse(templ1))

var reqBody2 = template.Must(template.New("request body 2").
	Parse(templ2))

var reqBody3 = template.Must(template.New("request body 3").
	Parse(templ3))

type Params struct {
	ID   string
	Name string
	Age  int
}

func main() {
	p := Params{"1", "mike", 10}

	// replace .ID and .Name
	if err := reqBody1.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
	// replace .Age
	if err := reqBody2.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
	// not replace .Unknown
	if err := reqBody3.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}
