package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nathanielfernandes/stars/lib/github"
)

func main() {
	m := github.NewManager()

	router := httprouter.New()
	router.GET("/:username/list", m.GetAllList)
	router.GET("/:username/map", m.GetAllMap)
	router.GET("/:username/langs", m.GetUsedLanguages)
	router.GET("/:username/image", m.GetImage)

	router.GET("/:username/repos/:repo", m.Get)

	fmt.Printf("stars\nListening on port 80\n")
	if err := http.ListenAndServe("0.0.0.0:80", router); err != nil {
		log.Fatal(err)
	}
}
