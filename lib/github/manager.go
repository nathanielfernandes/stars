package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nathanielfernandes/stars/ratelimit"
)

type Manager struct {
	c     http.Client
	Cache Repos
	Rlm   *ratelimit.RatelimitManager
}

func NewManager() Manager {
	return Manager{c: http.Client{}, Cache: Repos{}, Rlm: ratelimit.NewRatelimitManager(1, 60000)}
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
	data, _ := json.Marshal(m.Cache)
	w.Write(data)
}
