package main

import (
	"fmt"
	"net/http"
)

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)

}

func main() {
	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}
}
