package main

import (
	"fmt"
	"net/http"
	"url_shortner/handler"
)

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World: %s", r.URL.Path)
}

func main() {
	// create server mutex
	mux := newServeMux()

	// mapping handler to mutex
	pathMap := map[string]string{
		"/gg": "https://www.google.com/",
		"/go": "https://www.golang.org/",
	}

	mappedHandler := handler.MapHandler(pathMap, mux)

	// serving by using handler
	fmt.Println("start to serve at Port 8000...")
	http.ListenAndServe(":8000", mappedHandler)
}
