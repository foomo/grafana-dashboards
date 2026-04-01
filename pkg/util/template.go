package util

import (
	"strings"
	"text/template"
)

func MustRenderTemplate(v string, data any) string {
	r, err := RenderTemplate(v, data)
	if err != nil {
		panic(err)
	}
	return r
}

func RenderTemplate(v string, data any) (string, error) {
	// Create a new template and parse the string into it
	tmpl, err := template.New("").Parse(v)
	if err != nil {
		return "", err
	}

	// Execute the template, passing a string variable to replace the placeholder
	buf := new(strings.Builder)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
