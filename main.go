package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/nginx",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "nginx loadbalancer")
		},
	)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("error on starting server: %v", err)
	}
}
