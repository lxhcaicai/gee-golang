package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	for key, value := range req.Header {
		fmt.Fprintf(w, "Header[%q}] = %q\n", key, value)
	}
}
