package main

import (
	"html/template"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("param1"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param1")
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", param1)
}

func main3() {
	http.HandleFunc("/", handler2)
	http.ListenAndServe(":8080", nil)
}
