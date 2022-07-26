package github

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
)

var /* const */ AUTH_TOKEN = os.Getenv("github_token")
var /* const */ USER = os.Getenv("username")

func FetchRepos(c *http.Client, user string) (UserRepos, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+user+"/repos?per_page=100&visibility=public", nil)
	if err != nil {
		return UserRepos{}, nil
	}

	req.SetBasicAuth(USER, AUTH_TOKEN)

	res, err := c.Do(req)
	if err != nil {
		return UserRepos{}, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data UserRepos
	err = decoder.Decode(&data)

	if err != nil {
		return UserRepos{}, err
	}

	return data, nil
}

func FetchLangauges(c *http.Client, lang_url string) ([]Language, error) {
	req, err := http.NewRequest("GET", lang_url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(USER, AUTH_TOKEN)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data map[string]int
	err = decoder.Decode(&data)

	if err != nil {
		return nil, err
	}

	langs := make([]Language, 0, len(data))
	for lang, size := range data {
		langs = append(langs, Language{Name: lang, Size: size})
	}
	sort.Slice(langs, func(i, j int) bool { return langs[i].Size > langs[j].Size })

	return langs, nil
}

func FetchReadme(c *http.Client, repo string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+USER+"/"+repo+"/readme", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(USER, AUTH_TOKEN)

	res, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data Readme
	err = decoder.Decode(&data)

	if err != nil {
		return "", err
	}

	return data.Content, nil
}
