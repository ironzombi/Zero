package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	serve "zero/cmd"
)

type DataListener struct {
	input  string
	output string
	size   int
}

func (dl DataListener) Input(w http.ResponseWriter, req *http.Request) {
	dl.input = req.URL.Query().Get("input")
	fmt.Fprintf(w, "%s", dl.input)
}

func (dl DataListener) Echo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this worked,  %s", req.URL)
}

func (dl DataListener) Default(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}

// using a template, wont return anything  DataListener will be nil/nil
func (dl DataListener) Template(w http.ResponseWriter, req *http.Request) {
	tmplt := template.New("Hello World")
	tmplt, _ = tmplt.Parse("Top Secret: {{.Id}} - {{.Name}}!")
	p := DataListener{size: 1, output: "secret"}
	tmplt.Execute(w, p)
}

func main() {
	var dl DataListener

	http.HandleFunc("/", dl.Default)
	http.HandleFunc("/add", dl.Input)
	http.HandleFunc("/echo", dl.Echo)
	http.HandleFunc("/template", dl.Template)
	log.Fatal(http.ListenAndServe("192.168.1.117:8181", nil))
	serve.ZeroListener()
}
