package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

type Persons struct {
	Emails []Person
}

const tmpl = `
{{define "Emails"}}this is template T1: 
{{- range .}} 
	{{$name := .Name}}{{range .Emails}}Name is {{$name}}, email is {{.}} 
	{{end}}
{{end -}}
{{- end -}}
{{- define "ALL"}}{{template "Emails" .}}{{end -}}
{{define "Test1"}}{{23 -}} < {{- 45}}{{end}}
`

func main() {
	person := Person{
		Name:   "Satish",
		Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
	}

	t := template.New("Person")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.ExecuteTemplate(os.Stdout, "ALL", []Person{person})
	if err != nil {
		log.Fatal("ExecuteTemplate_ALL: ", err)
		return
	}

	err = t.ExecuteTemplate(os.Stdout, "Test1", []Person{person})
	if err != nil {
		log.Fatal("ExecuteTemplate_Temp1: ", err)
		return
	}

}
