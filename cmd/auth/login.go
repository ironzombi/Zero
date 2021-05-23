package login

import (
	"fmt"
	"net/http"
	"text/template"
)

/* while this is a login function it is more just
 * to work with input via html page served with golang */
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if checkCreds(r) {
			fmt.Println("username entered:", r.Form["username"])
			fmt.Println("password submitted:", r.Form["password"])
		} else {
			f, _ := template.ParseFiles("block.html")
			f.Execute(w, nil)
		}
	}
}

// TODO: creat an actual password check maybe.
func checkCreds(creds *http.Request) bool {

	if len(creds.Form["password"][0]) == 0 {
		fmt.Println("Password required")
		return false
	} else {
		fmt.Println("good boy")
		return true
	}

}
