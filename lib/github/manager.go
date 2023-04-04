package github

import (
	"fmt"
	"net/http"
	"net/url"
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
	allowed_users := strings.Split(os.Getenv("allow_list"), ",")
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

func GetFilterOptions(q url.Values) (bool, map[string]bool, map[string]bool) {
	allow_forks := q.Get("allow_forks")
	allow_forks_bool := false
	if allow_forks == "true" {
		allow_forks_bool = true
	}

	excluded_repos := q.Get("excluded_repos")
	excluded_repos_map := make(map[string]bool)
	if excluded_repos != "" {
		for _, repo := range strings.Split(excluded_repos, ",") {
			excluded_repos_map[strings.TrimSpace(repo)] = true
		}
	}

	excluded_languages := q.Get("excluded_languages")
	excluded_languages_map := make(map[string]bool)
	if excluded_languages != "" {
		for _, language := range strings.Split(excluded_languages, ",") {
			excluded_languages_map[strings.TrimSpace(language)] = true
		}
	}

	return allow_forks_bool, excluded_repos_map, excluded_languages_map
}

func (m *Manager) GetUsedLanguages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("username")
	if _, ok := ALLOWED_USERS[user]; ok {
		m.CheckUpdate(user)
		if rep, ok := m.Cache[user]; ok {
			allow_forks, excluded_repos, excluded_languages := GetFilterOptions(r.URL.Query())
			writeJson(w, GetUsedLanguages(rep.List, allow_forks, excluded_repos, excluded_languages))
		} else {
			writeJson(w, nil)
		}
	} else {
		writeJson(w, UNALLOWED_USER)
	}
}

func GetColorOptions(q url.Values) (string, string, string, string) {
	hashtag := func(s string) string {
		if s[0] != '#' {
			return "#" + s
		}
		return s
	}

	bgcolor := q.Get("bgcolor")
	if bgcolor == "" {
		bgcolor = "#1e1e1e"
	}

	outline := q.Get("outline")
	if outline == "" {
		outline = "#00d9ff"
	}

	textcolor := q.Get("textcolor")
	if textcolor == "" {
		textcolor = "#ffffffc4"
	}

	titlecolor := q.Get("titlecolor")
	if titlecolor == "" {
		titlecolor = "#ffffff"
	}

	return hashtag(bgcolor), hashtag(outline), hashtag(textcolor), hashtag(titlecolor)
}

func (m *Manager) GetImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := p.ByName("username")
	if _, ok := ALLOWED_USERS[user]; ok {
		m.CheckUpdate(user)
		if rep, ok := m.Cache[user]; ok {
			q := r.URL.Query()
			allow_forks, excluded_repos, excluded_languages := GetFilterOptions(q)
			languages := GetUsedLanguages(rep.List, allow_forks, excluded_repos, excluded_languages)

			bgcolor, outline, textcolor, titlecolor := GetColorOptions(q)
			im, err := GenImage(&m.c, languages, bgcolor, outline, textcolor, titlecolor)

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
			} else {
				imageResponse(w, im)
			}
		} else {
			fmt.Println("404")
			w.WriteHeader(404)
		}
	} else {
		writeJson(w, UNALLOWED_USER)
	}
}
