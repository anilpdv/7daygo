package main

import (
	"fmt"
	"log"
	"net/http"
)

//Engine : struct
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found %s\n", r.URL)
	}
}

func main() {
	engine := new(Engine)
	err := http.ListenAndServe(":8888", engine)
	if err != nil {
		log.Fatal(err)
	}
}
