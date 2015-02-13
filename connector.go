package invoicegenerator

import ()

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
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Unit_cost int    `json:"unit_cost"`
}

func (i *Invoice) Create() {
	b, err := json.M

}
