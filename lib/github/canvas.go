package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RunPayload struct {
	Size   []int   `json:"size"`
	Files  []File  `json:"files"`
	Assets []Asset `json:"assets"`
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
	dat, err := os.ReadFile("./quilt/langs.rv")
	if err != nil {
		panic(err)
	}
	return string(dat)
}()

func GenImage(c *http.Client, langs []Language, bgcolor, outline, textcolor, titlecolor string) (*bytes.Buffer, error) {
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

	payload := RunPayload{
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

	jsonBytes, err := json.Marshal(payload)
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
		fmt.Println(res.StatusCode)
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	return buf, nil
}
