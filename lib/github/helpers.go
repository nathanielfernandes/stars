package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	rl "github.com/nathanielfernandes/rl"
)

var repo_rlm = rl.NewRatelimitManager(1, 1000*60*60)

func (m *Manager) updateCache(user string) {
	repos, err := FetchRepos(&m.c, user)
	fmt.Println("FRESH GET")
	if err == nil {
		m.Cache[user] = repos.ToRepos(&m.c)
		fmt.Println("GOT")
	} else {
		fmt.Println(err)
	}
}

func (m *Manager) CheckUpdate(user string) {
	if !repo_rlm.IsRatelimited(user) {
		if !ALLOWED_USERS[user] {
			m.updateCache(user)
			ALLOWED_USERS[user] = true
		} else {
			go m.updateCache(user)
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
