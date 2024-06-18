package runescape

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	VALID_ALPHA = "abcdefghijklmnopqrstuvwxyz1234567890!&()/-+. "
)

// define struct to contain client and url
type Runescape struct {
	client *http.Client
	url    url.URL
}

// define struct to contain extra json from api response
type Valuation struct {
	Trend string `json:"trend"`
	Price int    `json:"price"`
}

// define main struct to contain all response json
// ordered based on relevance to primary endpoints
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Members     string `json:"members"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	IconLarge   string `json:"icon_large"`
	TypeIcon    string `json:"typeIcon"`

	Current Valuation `json:"current"`
	Today   Valuation `json:"today"`
}

// define struct to contain official api response
type ItemResponse struct {
	Total int    `json:"total"`
	Items []Item `json:"items"`
}

func New() *Runescape {
	/*
		Returns pointer to Runescape struct populated with relevant client and
		url infomation.
	*/

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	return &Runescape{
		client: &client,
		url: url.URL{
			Scheme: "https",
			Host:   "services.runescape.com",
			Path:   "/m=itemdb_rs/api/catalogue/items.json",
		},
	}
}

// create struct method
func (r *Runescape) GetItems(
	category int,
	alpha string,
	page int,
) (items []Item, err error) {
	/*
		Builds api url with arguments, category, alphabet, page and returns
		slice of Item from the response.
	*/
	queryParams := url.Values{}
	queryParams.Add("category", fmt.Sprintf("%d", category))
	queryParams.Add("alpha", alpha)
	queryParams.Add("page", fmt.Sprintf("%d", page))
	r.url.RawQuery = queryParams.Encode()

	resp, err := r.client.Get(r.url.String())
	if err != nil {
		return items, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return items, err
	}

	fmt.Println(string(body))

	var itemResponse ItemResponse
	err = json.Unmarshal(body, &itemResponse)
	if err != nil {
		return items, err
	}

	return itemResponse.Items, nil

}
