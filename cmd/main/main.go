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

	fmt.Printf("stars\nListening on port 80\n")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatal(err)
	}
}
