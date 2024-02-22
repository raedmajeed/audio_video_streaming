package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const directory = "files"
	const port = 8080

	http.Handle("/", addHeaders(http.FileServer(http.Dir(directory))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", directory, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
