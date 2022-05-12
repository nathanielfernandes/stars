# stars

A little service that I use on my [personal site](https://nathanielfernandes.ca) to get live repo stats without worrying about vistors being ratelimited by github.

#### usage

set the env var

```
username=<your github username>
github_token=<your github auth token>
```

run with docker

#### endpoints

theres just one endpoint that returns repo stats -> `/`

###### example response (JSON)

```json
{
  "lingo": {
    "stars": 38,
    "tags": ["charm", "cli", "golang", "tokei"],
    "forks": 1,
    "created": "2022-02-23T01:04:53Z"
  },
  "HamoodBot": {
    "stars": 13,
    "tags": ["chess", "discord-bot", "hamood", "math", "meme"],
    "forks": 2,
    "created": "2020-07-06T20:39:02Z"
  },
  "Rustacean-Tracing": {
    "stars": 7,
    "tags": [],
    "forks": 0,
    "created": "2021-08-21T00:53:34Z"
  },
  "chirp": {
    "stars": 6,
    "tags": ["chip8", "emulator", "macroquad"],
    "forks": 0,
    "created": "2022-04-21T05:12:27Z"
  },

  "mime": {
    "stars": 11,
    "tags": ["meme-api", "meme-generator", "memes"],
    "forks": 0,
    "created": "2022-03-27T04:18:13Z"
  }
}
```
