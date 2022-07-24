package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	rl "github.com/nathanielfernandes/rl"
)

var EMPTY bool = true
var repo_rlm = rl.NewRatelimitManager(1, 1000*60*60)

func (m *Manager) updateCache() {
	repos, err := FetchRepos(&m.c)
	fmt.Println("FRESH GET")
	if err == nil {
		m.Cache = repos.ToRepos(&m.c)
		fmt.Println("GOT")
	} else {
		fmt.Println(err)
	}
}

func (m *Manager) CheckUpdate() {
	if !repo_rlm.IsRatelimited("GENERIC") {
		if EMPTY {
			m.updateCache()
			EMPTY = false
		} else {
			go m.updateCache()
		}
	}
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func writeJson(w http.ResponseWriter, data interface{}) {
	addHeaders(w)
	json, _ := json.MarshalIndent(data, "", "   ")
	w.Write(json)
}
