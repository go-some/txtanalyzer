package txtanalyzer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func filterEntities(entList []Entity) ([]string, []string, []string) {
	var personList, orgList, prodList []string
	for _, ent := range entList {
		txt := strings.ToLower(strings.Trim(ent.Text, " "))
		if ent.Label == "PERSON" {
			personList = append(personList, txt)
		} else if ent.Label == "ORG" {
			orgList = append(orgList, txt)
		} else if ent.Label == "PRODUCT" {
			prodList = append(prodList, txt)
		}
	}
	return personList, orgList, prodList
}

func NEROnDoc(title, body string) ([]Entity, []string, []string, []string) {
	entitiesInTitle, nerErr1 := RequestNER(title)
	if nerErr1 != nil {
		log.Fatal(nerErr1)
	}
	entitiesInBody, nerErr2 := RequestNER(body)
	if nerErr2 != nil {
		log.Fatal(nerErr2)
	}
	personList, orgList, prodList := filterEntities(entitiesInBody)
	return entitiesInTitle, personList, orgList, prodList
}
