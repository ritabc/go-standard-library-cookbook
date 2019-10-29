// Communictate with HTTP server at higher level. Refer to 7.02 for connection with TCP server at a lower level
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello World!"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	simplePost()
	useRequest()
}

// A common way to connect to HTTP server
func simplePost() {
	res, err := http.Post("http://localhost:7070", "application/x-www-form-urlencoded", strings.NewReader("name=rita&lastname=bc"))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println("Response from server: ", string(data))
}

// Another common way to connect to HTTP server, but with a more customizable API and its own instance of Client
func useRequest() {
	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "RITA")
	form.Add("lastname", "BC")

	req, err := http.NewRequest("POST", "http://localhost:7070", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println("Response from server: ", string(data))
}
