package github

import (
	"net/http"
)

type RespRepo struct {
	// Private      bool     `json:"private"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Stars        int      `json:"stargazers_count"`
	Tags         []string `json:"topics"`
	Forks        int      `json:"forks_count"`
	Created      string   `json:"created_at"`
	Page         string   `json:"homepage"`
	LanguagesUrl string   `json:"languages_url"`
}

type Readme struct {
	Content string `json:"content"`
}

type UserRepos []RespRepo

type Repo struct {
	Stars       int      `json:"stars"`
	Tags        []string `json:"tags"`
	Forks       int      `json:"forks"`
	Created     string   `json:"created"`
	Description string   `json:"description"`
	Page        string   `json:"page"`
	Languages   []string `json:"languages"`
	// Image       string   `json:"image"`
}

type Repos map[string]Repo

func (r RespRepo) ToData() Repo {
	return Repo{Stars: r.Stars, Tags: r.Tags, Forks: r.Forks, Created: r.Created, Description: r.Description, Page: r.Page}
}

func (ur UserRepos) ToRepos(c *http.Client) Repos {
	m := make(map[string]Repo)
	for _, repo := range ur {
		r := repo.ToData()
		if langs, err := FetchLangauges(c, repo.Name); err == nil {
			r.Languages = langs
		}
		m[repo.Name] = r
	}

	return m
}
