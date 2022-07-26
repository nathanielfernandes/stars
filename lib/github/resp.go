package github

import (
	"net/http"
	"sort"
	"sync"
	"time"
)

type RespRepo struct {
	Name         string   `json:"name"`
	IsFork       bool     `json:"fork"`
	Description  string   `json:"description"`
	Stars        int      `json:"stargazers_count"`
	Tags         []string `json:"topics"`
	Forks        int      `json:"forks_count"`
	Created      string   `json:"created_at"`
	Updated      string   `json:"pushed_at"`
	Page         string   `json:"homepage"`
	LanguagesUrl string   `json:"languages_url"`
}

type Readme struct {
	Content string `json:"content"`
}

type UserRepos []RespRepo

type Language struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Repo struct {
	Name        string     `json:"name,omitempty"`
	Stars       int        `json:"stars"`
	Tags        []string   `json:"tags"`
	Forks       int        `json:"forks"`
	Created     int64      `json:"created"`
	Updated     int64      `json:"updated"`
	Description string     `json:"description"`
	Page        string     `json:"page"`
	Languages   []Language `json:"languages"`
	IsFork      bool       `json:"is_fork"`
}

type Repos struct {
	Map  map[string]Repo
	List []Repo
}

func (r RespRepo) ToData() Repo {
	created, _ := time.Parse("2006-01-02T15:04:05Z", r.Created)
	updated, _ := time.Parse("2006-01-02T15:04:05Z", r.Updated)
	return Repo{Stars: r.Stars, Tags: r.Tags, Forks: r.Forks, Created: created.UnixMilli(), Updated: updated.UnixMilli(), Description: r.Description, Page: r.Page, IsFork: r.IsFork}
}

func addLangs(c *http.Client, wg *sync.WaitGroup, url string, name string, index int, m map[string]Repo, l []Repo) {
	defer wg.Done()

	if langs, err := FetchLangauges(c, url); err == nil {
		r := m[name]
		r.Languages = langs
		m[name] = r
		l[index].Languages = langs
	}
}

func (ur UserRepos) ToRepos(c *http.Client) Repos {
	m := make(map[string]Repo)
	l := make([]Repo, 0, len(ur))

	wg := &sync.WaitGroup{}
	wg.Add(len(ur))

	for i, repo := range ur {
		r := repo.ToData()

		m[repo.Name] = r
		r.Name = repo.Name
		l = append(l, r)

		go addLangs(c, wg, repo.LanguagesUrl, repo.Name, i, m, l)
	}

	wg.Wait()

	sort.Slice(l, func(i, j int) bool { return l[i].Updated > l[j].Updated })

	return Repos{Map: m, List: l}
}
