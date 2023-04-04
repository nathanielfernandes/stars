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

func addLangs(c *http.Client, wg *sync.WaitGroup, lock *sync.Mutex, url string, name string, index int, m map[string]Repo, l []Repo, langs map[string]int) {
	defer wg.Done()

	if ls, err := FetchLangauges(c, url); err == nil {
		lock.Lock()

		r := m[name]
		r.Languages = ls
		m[name] = r
		l[index].Languages = ls

		// if !l[index].IsFork {
		for _, lang := range ls {
			langs[lang.Name] += lang.Size
		}
		// }

		lock.Unlock()
	}
}

func (ur UserRepos) ToRepos(c *http.Client) Repos {
	m := make(map[string]Repo)
	lock := sync.Mutex{}
	langs := make(map[string]int)
	l := make([]Repo, 0, len(ur))

	wg := &sync.WaitGroup{}
	wg.Add(len(ur))

	for i, repo := range ur {
		r := repo.ToData()

		m[repo.Name] = r
		r.Name = repo.Name
		l = append(l, r)

		go addLangs(c, wg, &lock, repo.LanguagesUrl, repo.Name, i, m, l, langs)
	}

	wg.Wait()

	sort.Slice(l, func(i, j int) bool { return l[i].Updated > l[j].Updated })

	languages := make([]Language, 0, len(langs))
	for name, size := range langs {
		languages = append(languages, Language{Name: name, Size: size})
	}

	sort.Slice(languages, func(i, j int) bool { return languages[i].Size > languages[j].Size })

	return Repos{Map: m, List: l}
}

func GetUsedLanguages(l []Repo, forks bool, exclude_repos map[string]bool, exclude_langs map[string]bool, threshold float64) []Language {
	langs := make(map[string]int)

	total := 0
	for _, r := range l {
		if exclude_repos[r.Name] {
			continue
		}

		if !forks && r.IsFork {
			continue
		}

		for _, lang := range r.Languages {
			if exclude_langs[lang.Name] {
				continue
			}

			langs[lang.Name] += lang.Size
			total += lang.Size
		}
	}

	languages := make([]Language, 0, len(langs))
	for name, size := range langs {
		if float64(size)/float64(total) < threshold {
			continue
		}

		languages = append(languages, Language{Name: name, Size: size})
	}

	sort.Slice(languages, func(i, j int) bool { return languages[i].Size > languages[j].Size })

	return languages
}
