package main

import (
	"fmt"
	"net/http"
)

const addr = "localhost:7070"

// RedirectServer is a server which will redirect upon a req
type RedirectServer struct {
	redirectCount int
}

func (s *RedirectServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.redirectCount++
	fmt.Println("Received header: ", req.Header.Get("Known-redirects"))
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func main() {
	s := http.Server{
		Addr:    addr,
		Handler: &RedirectServer{0},
	}
	go s.ListenAndServe()

	client := http.Client{}
	redirectCount := 0

	// If the count of redirects is reached, return error
	// If client has CheckRedirect func, client calls it before following an HTTP redirect
	// req is upcoming request, via is prior requests, oldest first
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		if redirectCount > 2 {
			return fmt.Errorf("Too many redirects")
		}
		req.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}
}
