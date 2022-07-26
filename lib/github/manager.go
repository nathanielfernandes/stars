package github

import (
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Manager struct {
	c     http.Client
	Cache map[string]Repos
}

func NewManager() Manager {
	return Manager{c: http.Client{}, Cache: map[string]Repos{}}
}

func getAllowed() map[string]bool {
	m := map[string]bool{}
	allowed_users := strings.Split(os.Getenv("allowed_users"), ",")
	for _, user := range allowed_users {
		m[user] = false
	}

	return m
}

var /* const */ ALLOWED_USERS = getAllowed()

var UNALLOWED_USER = "User not allowed!"

func (m *Manager) GetAllList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("username")
	if _, ok := ALLOWED_USERS[user]; ok {
		m.CheckUpdate(user)
		writeJson(w, m.Cache[user].List)
	} else {
		writeJson(w, UNALLOWED_USER)
	}
}

func (m *Manager) GetAllMap(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("username")
	if _, ok := ALLOWED_USERS[user]; ok {
		m.CheckUpdate(user)
		writeJson(w, m.Cache[user].Map)
	} else {
		writeJson(w, UNALLOWED_USER)
	}
}

func (m *Manager) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("username")
	if _, ok := ALLOWED_USERS[user]; ok {
		m.CheckUpdate(user)
		repo := p.ByName("repo")
		if r, ok := m.Cache[user].Map[repo]; ok {
			writeJson(w, r)
		} else {
			writeJson(w, nil)
		}
	} else {
		writeJson(w, UNALLOWED_USER)
	}
}
