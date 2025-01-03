package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RunPayload struct {
	Size   []int   `json:"size"`
	Files  []File  `json:"files"`
	Assets []Asset `json:"assets"`

	Frames int  `json:"frames"`
	Repeat int  `json:"repeat"`
	Looped bool `json:"looped"`
	Delay  int  `json:"delay"`
}

type File struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ImageAsset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LiteralAsset struct {
	Name    string `json:"name"`
	Literal string `json:"literal"`
}

type Asset interface{}

var /* const */ API = os.Getenv("api_endpoint")
var /* const */ SECRET = os.Getenv("canvas_secret")
var /* const */ LANGS_CODE = func() string {
	dat, err := os.ReadFile("./quilt/langs.ql")
	if err != nil {
		panic(err)
	}
	return string(dat)
}()

func GenPayload(c *http.Client, langs []Language, bgcolor, outline, textcolor, titlecolor string) RunPayload {
	total := 0.0
	for _, lang := range langs {
		total += float64(lang.Size)
	}

	literal := "["
	for _, lang := range langs {
		literal += fmt.Sprintf("[\"%s\",%f,%s],\n", lang.Name, float64(lang.Size)/total, GetColorFromAlias(lang.Name))
	}
	literal += "]"

	h := ((len(langs) - 1) / 2) * 25

	return RunPayload{
		Size: []int{300, 100 + h},
		Files: []File{
			{
				Name: "main",
				Code: LANGS_CODE,
			},
		},
		Assets: []Asset{
			LiteralAsset{
				Name:    "langs",
				Literal: literal,
			},
			LiteralAsset{
				Name:    "bgcolor",
				Literal: bgcolor,
			},
			LiteralAsset{
				Name:    "outline",
				Literal: outline,
			},
			LiteralAsset{
				Name:    "textcolor",
				Literal: textcolor,
			},
			LiteralAsset{
				Name:    "titlecolor",
				Literal: titlecolor,
			},
		},
	}
}

func GenImage(c *http.Client, langs []Language, bgcolor, outline, textcolor, titlecolor string) (*bytes.Buffer, error) {
	payload := GenPayload(c, langs, bgcolor, outline, textcolor, titlecolor)
	return payload.Run(c)
}

func GenGif(c *http.Client, langs []Language, bgcolor, outline, textcolor, titlecolor string) (*bytes.Buffer, error) {
	payload := GenPayload(c, langs, bgcolor, outline, textcolor, titlecolor)

	payload.Frames = 100
	payload.Repeat = 0
	payload.Delay = 20

	return payload.Run(c)
}

func (p *RunPayload) Run(c *http.Client) (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest("POST", API+"/run/"+SECRET, body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err // or log the error
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString) // Print the response body as a string
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	return buf, nil
}
