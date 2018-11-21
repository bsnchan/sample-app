package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it works!"))
	})

	log.Print("starting application...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
