package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello this is my first web-go %s\n", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
