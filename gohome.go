package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	paths := getPaths()
	for _, path := range paths {
		r.HandleFunc(path.Path, pathHandler)
	}

	fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
	http.Handle("/js/", fileServer)
	fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
