package github

type RespRepo struct {
	Private bool     `json:"private"`
	Name    string   `json:"name"`
	Stars   int      `json:"stargazers_count"`
	Tags    []string `json:"topics"`
	Forks   int      `json:"forks_count"`
	Created string   `json:"created_at"`
}

type UserRepos []RespRepo

type Repo struct {
	Stars   int      `json:"stars"`
	Tags    []string `json:"tags"`
	Forks   int      `json:"forks"`
	Created string   `json:"created"`
}

type Repos map[string]Repo

func (r RespRepo) ToData() Repo {
	return Repo{Stars: r.Stars, Tags: r.Tags, Forks: r.Forks, Created: r.Created}
}

func (ur UserRepos) ToRepos() Repos {
	m := make(map[string]Repo)
	for _, repo := range ur {
		if !repo.Private {
			m[repo.Name] = repo.ToData()
		}
	}

	return m
}
