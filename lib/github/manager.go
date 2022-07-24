package github

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Manager struct {
	c     http.Client
	Cache Repos
}

func NewManager() Manager {
	return Manager{c: http.Client{}, Cache: Repos{}}
}

func (m *Manager) GetAllList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m.CheckUpdate()
	writeJson(w, m.Cache.List)
}

func (m *Manager) GetAllMap(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m.CheckUpdate()
	writeJson(w, m.Cache.Map)
}

func (m *Manager) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m.CheckUpdate()
	repo := p.ByName("repo")
	if r, ok := m.Cache.Map[repo]; ok {
		writeJson(w, r)
	} else {
		writeJson(w, nil)
	}
}
