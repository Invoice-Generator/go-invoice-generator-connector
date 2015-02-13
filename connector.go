package invoicegenerator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const requestURL = "https://invoice-generator.com"
const requestType = "application/json"

type Invoice struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Logo   string `json:"logo"`
	Number int    `json:"number"`
	Date   string `json:"date"`
	Items  []Item `json:"items"`
	Notes  string `json:"notes"`
	Terms  string `json:"terms"`
}

type Item struct {
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	Unit_cost float64 `json:"unit_cost"`
}

func (i *Invoice) Create(filename string, fileDir string) {
	b, err := json.Marshal(*i)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Post(requestURL, requestType, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	f := fileDir + filename + ".pdf"

	err = ioutil.WriteFile(f, body, 0644)

	if err != nil {
		panic(err)
	}

}
