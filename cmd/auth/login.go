package login

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type PiHosts struct {
	addr     []string
	hostname []string
}

var Hosts PiHosts

/* while this is a login function it is more just
 * to work with input via html page served with golang */
func Login(w http.ResponseWriter, r *http.Request) {
	cTime := time.Now()
	fmt.Printf("%s method:%s\n", cTime.Format("2006.01.02 15:04:05"), r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("site/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if checkCreds(r) {
			fmt.Println("username entered:", r.Form["username"])
			fmt.Println("password entered:", r.Form["password"])
		} else {
			f, _ := template.ParseFiles("site/block.html")
			f.Execute(w, nil)
		}
	}
}

// TODO: this currently doesnt work, dl.Datalistener is empty when called.
func Render(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		t, _ := template.ParseFiles("site/info.html")
		t.Execute(w, nil)
	} else {
		req.ParseForm()
		Hosts.addr = req.Form["ipaddress"]
		Hosts.hostname = req.Form["hostname"]
		t, err := template.New("MIPHACK").Parse("Zero \"{{ .hostname}}\"")
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, Hosts)
	}
}

func ListServers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Raspberry Pi Zeros are awesome: %s", Hosts.hostname)
}

// TODO: creat an actual password check maybe.
func checkCreds(creds *http.Request) bool {

	if len(creds.Form["password"][0]) == 0 {
		fmt.Println("Password required")
		return false
	} else {
		fmt.Println("authenticated")
		return true
	}

}
