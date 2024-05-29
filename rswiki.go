package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// getItems retreives items from a JSON API endpoint
func getItems() (items []Item, err error) {
	url := "https://api.runeprofitforge.com/items"

	resp, err := http.Get(url)
	if err != nil {
		return items, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return items, nil
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return items, nil
	}

	return items, nil

}
