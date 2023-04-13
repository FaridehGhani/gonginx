package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hii from nginx")
		},
	)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("error on starting server: %v", err)
	}
}
