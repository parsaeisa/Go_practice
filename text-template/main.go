package main

import (
	"bytes"
	"strings"
	"text/template"
)

type Car struct {
	// for template , attributes must be exposable
	Lat float32
	Lng float32
}

func Example1() string {
	car := Car{
		Lat: 12.2,
		Lng: 13.2,
	}
	c, err := template.New("content-template").Parse("خودروی شما در موقعیت {{.Lat}} و {{.Lng}} قرار دارد .")
	if err != nil {
		panic(err)
	}

	// what is bytes buffer ?
	var tpl bytes.Buffer
	err = c.Execute(&tpl, car)
	if err != nil {
		return ""
	}

	content := tpl.String()
	return content
}

func Example2() string {

	var temp *template.Template = template.Must(
		template.New("temp").Parse("The name is {{.Name}} \n" +
			`{{ if .PrintHeight}}` +
			`The height is .Height {{ end }}`,
		))

	// A new way to define a struct with an instance of it
	a := struct {
		Name        string
		Height      float64
		PrintHeight bool
	}{
		Name:        "Parsa",
		Height:      1.92,
		PrintHeight: true,
	}

	var builder strings.Builder
	if err := temp.Execute(&builder, a); err != nil {
		panic(err)
	}

	return builder.String()
}

func main() {
	println(Example1())
	println("=================")
	print(Example2())
}
