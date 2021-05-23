package serve

import (
	"fmt"
	"log"
	"net/http"
	login "zero/cmd/auth"
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

func ZeroListener() {
	var dl DataListener
	http.HandleFunc("/", dl.Default)
	http.HandleFunc("/add", dl.Input)
	http.HandleFunc("/echo", dl.Echo)
	http.HandleFunc("/template", login.Render)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/info", login.ListServers)

	log.Fatal(http.ListenAndServe("0.0.0.0:8181", nil))

}
