package main

import (
	"bytes"
	"text/template"
)

type Car struct {
	// for template , attributes must be exposable
	Lat float32
	Lng float32
}

func main() {
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
		return
	}

	content := tpl.String()
	println(content)
}
