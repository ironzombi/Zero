package serve

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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

// TODO: this currently doesnt work, dl.Datalistener is empty when called.
func (dl *DataListener) Template(w http.ResponseWriter, req *http.Request) {

	t, err := template.New("MIPHACK").Parse("Input \"{{ .dl.input}}\"")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, dl)
}

func ZeroListener() {
	var dl DataListener
	http.HandleFunc("/", dl.Default)
	http.HandleFunc("/add", dl.Input)
	http.HandleFunc("/echo", dl.Echo)
	http.HandleFunc("/template", dl.Template)

	log.Fatal(http.ListenAndServe("192.168.1.117:8181", nil))
	fmt.Println("check good")
}
