package main

import (
	"fmt"
	"html/template"
)

type Templates struct {
	T []string
}

func ParseTemplates(templates ...string) (*template.Template, error) {
	return parseTemplates(nil, templates...)
}

func parseTemplates(t *template.Template, templates ...string) (*template.Template, error) {
	if len(templates) == 0 {
		return nil, fmt.Errorf("No templates specified in call to ParseTemplates.")
	}
	name := "temporary.template.name"
	for _, tv := range templates {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		tmpl = t

		_, err := tmpl.Parse(tv)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
