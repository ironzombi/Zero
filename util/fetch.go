package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// download url as a string
func FetchPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", body)
}
