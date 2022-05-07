package github

type RespRepo struct {
	Private bool     `json:"private"`
	Name    string   `json:"name"`
	Stars   int      `json:"stargazers_count"`
	Tags    []string `json:"topics"`
}

type UserRepos []RespRepo

type Repo struct {
	Stars int      `json:"stars"`
	Tags  []string `json:"tags"`
}

type Repos map[string]Repo

func (r RespRepo) ToData() Repo {
	return Repo{Stars: r.Stars, Tags: r.Tags}
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
