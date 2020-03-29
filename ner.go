package txtanalyzer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Entity struct {
	Label string
	Text  string
}

type ResNER struct {
	Ent_list []Entity
	Status   int
}

func RequestNER(text string) ([]Entity, error) {
	API_ADDR := "http://localhost:5000/ner"
	resp, err := http.PostForm(API_ADDR, url.Values{"text": {text}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Entity{}, err
	}

	var resNer ResNER
	json.Unmarshal(respBody, &resNer)

	return resNer.Ent_list, nil
}
