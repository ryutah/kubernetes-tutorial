package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Call root")
		fmt.Fprintln(w, "Call in app1")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Call health")
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}
