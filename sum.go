package txtanalyzer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ResSUM struct {
	Sum_text string
	Status   int
}

func RequestSUM(text string) (string, error) {
	API_ADDR := "http://localhost:5000/sum"
	resp, err := http.PostForm(API_ADDR, url.Values{"text": {text}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var resSum ResSUM
	json.Unmarshal(respBody, &resSum)

	return resSum.Sum_text, nil
}
