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
	Rlm   *rl.RatelimitManager
}

func NewManager() Manager {
	return Manager{c: http.Client{}, Cache: Repos{}, Rlm: rl.NewRatelimitManager(1, 60000)}
}

func (m *Manager) Get(w http.ResponseWriter, r *http.Request) {
	if !m.Rlm.IsRatelimited("GENERIC") {
		repos, err := FetchRepos(&m.c)
		fmt.Println("FRESH GET")
		if err == nil {
			m.Cache = repos.ToRepos()
			fmt.Println("GOT")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, _ := json.Marshal(m.Cache)
	w.Write(data)
}
