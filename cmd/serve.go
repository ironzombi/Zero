package serve

import (
	"fmt"
	"log"
	"net/http"
	login "zero/cmd/auth"
	plot "zero/cmd/plot"
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
	fmt.Fprintf(w, "repeats,  %s", req.URL)
}

func (dl DataListener) Default(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "site/index.html")
}

func ZeroListener() {
	var dl DataListener
	http.HandleFunc("/", dl.Default)
	http.HandleFunc("/add", dl.Input)
	http.HandleFunc("/echo", dl.Echo)
	http.HandleFunc("/template", login.Render)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/info", login.ListServers)
	http.HandleFunc("/plot", plot.Plot)
	log.Fatal(http.ListenAndServe("0.0.0.0:8181", nil))

}
