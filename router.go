package main

import (
	"fmt"
)

type Path struct {
	Path string "path"
}

func pathHandler(w http.ResponseWriter, req *http.Request) {

	templates := &Templates{[]string{
		`{{define "Items"}}<div><ul>{{range $v := .Paths}}<li><a href="http://localhost:8080{{$v.Path}}">{{$v.Path}}</a></li>{{end}}</ul></div>{{end}}`,
		`{{define "Content"}}<div>{{.Path}}</div><div>{{template "Items" .}}</div>{{end}}`,
		`<html><head><title>{{.Path}}</title></head><body>{{template "Content" .}}</body></html>`,
	}}

	var data = make(map[string]interface{})
	data["Path"] = req.URL.Path
	fmt.Println(data["Path"])
	paths := getPaths()
	data["Paths"] = paths
	t, err := ParseTemplates(templates.T...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	var data = make(map[string]string)
	data["scheme"] = req.URL.Scheme
	data["host"] = req.URL.Host
	data["path"] = req.URL.Path
	fmt.Println(w.Header())
	templates := &Templates{[]string{
		`{{define "Content"}}<h1>404</h1><div>{{.scheme}}://{{.host}}{{.path}} was not found</div>{{end}}`,
		`<html><head><title>{{.scheme}}://{{.host}}{{.path}}</title></head><body>{{template "Content" .}}</body></html>`,
	}}
	t, err := ParseTemplates(templates.T...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func getPaths() []Path {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("mux")

	result := []Path{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}
