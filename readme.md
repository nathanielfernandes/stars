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


`GET /list` (sorted by last updated)
```json
[
   {
      "name": "stars",
      "stars": 1,
      "tags": [],
      "forks": 1,
      "created": 1651899768, // Unix Seconds
      "updated": 1658610278, // Unix Seconds
      "description": "A little service that I use on my personal site to get live repo stats without worrying about vistors being ratelimited by github.",
      "page": "https://stars.ncp.nathanferns.xyz/",
      "languages": [
         {
            "name": "Go",
            "size": 4540 // bytes
         },
         {
            "name": "Dockerfile",
            "size": 186 // bytes
         },
         {
            "name": "Makefile",
            "size": 46 // bytes
         }
      ],
      "is_fork": false
   },
   ...
]
```

`GET /map`
```json
{
   "stars": {
      "stars": 1,
      "tags": [],
      "forks": 1,
      "created": 1651899768, // Unix Seconds
      "updated": 1658610278, // Unix Seconds
      "description": "A little service that I use on my personal site to get live repo stats without worrying about vistors being ratelimited by github.",
      "page": "https://stars.ncp.nathanferns.xyz/",
      "languages": [
         {
            "name": "Go",
            "size": 4540 // bytes
         },
         {
            "name": "Dockerfile",
            "size": 186 // bytes
         },
         {
            "name": "Makefile",
            "size": 46 // bytes
         }
      ],
      "is_fork": false
   },
   ...
}
```

`GET /repos/:name`
```json
{
  "stars": 1,
  "tags": [],
  "forks": 1,
  "created": 1651899768, // Unix Seconds
  "updated": 1658610278, // Unix Seconds
  "description": "A little service that I use on my personal site to get live repo stats without worrying about vistors being ratelimited by github.",
  "page": "https://stars.ncp.nathanferns.xyz/",
  "languages": [
      {
        "name": "Go",
        "size": 4540 // bytes
      },
      {
        "name": "Dockerfile",
        "size": 186 // bytes
      },
      {
        "name": "Makefile",
        "size": 46 // bytes
      }
  ],
  "is_fork": false
}
```