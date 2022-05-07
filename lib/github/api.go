package github

import (
	"encoding/json"
	"net/http"
	"os"
)

var /* const */ AUTH_TOKEN = os.Getenv("github_token")

const USER = "nathanielfernandes"
const ENDPOINT = "https://api.github.com/users/" + USER + "/repos?per_page=100"

func FetchRepos(c *http.Client) (UserRepos, error) {
	req, err := http.NewRequest("GET", ENDPOINT, nil)
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
