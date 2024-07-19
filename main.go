package main

import (
	"fmt"
	"net/http"
	"os"
)

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	if req.URL.Path == "/index.html" {
		res.Header().Add("Content-Type", "text/html")
		content, err := os.ReadFile("index.html")
		if err == nil {
			res.Write(content)
		} else {
		}
	} else if req.URL.Path == "/clicked" {
		res.Header().Add("Content-Type", "text/html")
		res.Write([]byte("Nice!"))
	}
}

func main() {
	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}
}
