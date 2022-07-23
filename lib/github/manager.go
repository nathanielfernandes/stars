package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	rl "github.com/nathanielfernandes/rl"
)

type Manager struct {
	c     http.Client
	Cache Repos
}

func NewManager() Manager {
	return Manager{c: http.Client{}, Cache: Repos{}}
}

var EMPTY bool = true

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

var repo_rlm = rl.NewRatelimitManager(1, 1000*60*60)

func (m *Manager) Get(w http.ResponseWriter, r *http.Request) {
	if !repo_rlm.IsRatelimited("GENERIC") {
		if EMPTY {
			m.updateCache()
			EMPTY = false
		} else {
			go m.updateCache()
		}
		fmt.Println("Not RATE LIMITED")
	} else {
		fmt.Println("RATE LIMITED")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, _ := json.MarshalIndent(m.Cache, "", "   ")
	w.Write(data)
}
