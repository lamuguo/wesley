package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	httpAddr = flag.String("http", "0.0.0.0:8080", "HTTP listen address")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("starting server...\n")
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
