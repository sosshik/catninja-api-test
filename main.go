package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

type Response struct {
	Data []Cat `json:"data"`
}

type Cat struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

func main() {
	num := flag.Int("limit", 100, "limit of cats")
	flag.Parse()
	limit := *num

	url := fmt.Sprintf("https://catfact.ninja/breeds?limit=%d", limit)

	body, err := getJson(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cats, err := jsonToCats(body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	catsByCountry := make(map[string][]Cat)

	for _, cat := range cats {
		if cat.Country == "" {
			catsByCountry["Undefined"] = append(catsByCountry["Undefined"], cat)
			continue
		}
		catsByCountry[cat.Country] = append(catsByCountry[cat.Country], cat)
	}

	sortByCoat(catsByCountry)

	jsonOut, err := json.Marshal(catsByCountry)

	if err != nil {
		fmt.Printf("Unable to write: %v\n", err)
		os.Exit(1)
	}

	ioutil.WriteFile("out.json", jsonOut, 0644)

}

func getJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to make GET request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to get response body: %v", err)
	}
	return body, nil
}

func jsonToCats(body []byte) ([]Cat, error) {

	respObj := Response{}

	err := json.Unmarshal(body, &respObj)
	if err != nil {
		return []Cat{}, fmt.Errorf("unable to unmarshal JSON: %v", err)
	}

	return respObj.Data, nil
}

func sortByCoat(catsByCountry map[string][]Cat) {
	hairLength := map[string]int{
		"":                    0,
		"Hairless":            1,
		"Partly Hairless":     2,
		"Hairless/Furry down": 3,
		"Short":               4,
		"Short/Hairless":      5,
		"Medium":              6,
		"Semi Long":           7,
		"Semi-long":           8,
		"Short/Long":          9,
		"Long/Short":          10,
		"Long/short":          11,
		"Rex (Short/Long)":    12,
		"Rex":                 13,
		"Long":                14,
		"All":                 15,
	}

	for _, cats := range catsByCountry {
		sort.Slice(cats, func(i, j int) bool {
			return hairLength[cats[i].Coat] > hairLength[cats[j].Coat]
		})
	}
}
