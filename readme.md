# stars

A little service that I use on my [personal site](https://nathanielfernandes.ca) to get live repo stats without worrying about vistors being ratelimited by github.

#### usage

set the env var

```
github_token=<your github auth token>
```

edit your username in `./lib/github/api.go`

run with docker

#### endpoints

theres just one endpoint that returns repo stats -> `/`
