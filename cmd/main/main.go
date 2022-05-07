package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanielfernandes/stars/lib/github"
)

func main() {
	m := github.NewManager()
	http.HandleFunc("/", m.Get)

	fmt.Printf("Go Get Some\nListening on port 80\n")
	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		log.Fatal(err)
	}
}
