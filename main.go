package main

import (
	"fmt"
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

func main() {
	var dl DataListener

	http.HandleFunc("/", dl.Default)
	http.HandleFunc("/add", dl.Input)
	http.HandleFunc("/echo", dl.Echo)
	log.Fatal(http.ListenAndServe("192.168.1.117:8181", nil))
	serve.ZeroListener()
}
