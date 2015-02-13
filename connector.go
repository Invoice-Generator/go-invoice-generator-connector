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
	Header             string `json:"header,omitempty"`
	ToTitle            string `json:"to_title,omitempty"`
	InvoiceNumberTitle string `json:"invoice_number_title,omitempty"`
	DateTitle          string `json:"date_title,omitempty"`
	PaymentTermsTitle  string `json:"payment_terms_title,omitempty"`
	DueDateTitle       string `json:"due_date_title,omitempty"`
	PurchaseOrderTitle string `json:"purchase_order_title,omitempty"`
	QuantityHeader     string `json:"quantity_header,omitempty"`
	ItemHeader         string `json:"item_header,omitempty"`
	UnitCostHeader     string `json:"unit_cost_header,omitempty"`
	AmountHeader       string `json:"amount_header,omitempty"`
	SubtotalTitle      string `json:"subtotal_title,omitempty"`
	DiscountsTitle     string `json:"discounts_title,omitempty"`
	TaxTitle           string `json:"tax_title,omitempty"`
	ShippingTitle      string `json:"shipping_title,omitempty"`
	TotalTitle         string `json:"total_title,omitempty"`
	AmountPaidTitle    string `json:"amount_paid_title,omitempty"`
	BalanceTitle       string `json:"balance_title,omitempty"`
	TermsTitle         string `json:"terms_title,omitempty"`
	NotesTitle         string `json:"notes_title,omitempty"`

	Currency string `json:"currency,omitempty"`

	From          string  `json:"from"`
	To            string  `json:"to"`
	Logo          string  `json:"logo,omitempty"`
	Number        string  `json:"number,omitempty"`
	PurchaseOrder string  `json:"purchase_order,omitempty"`
	Date          string  `json:"date"`
	PaymentTerms  string  `json:"payment_terms,omitempty"`
	Items         []Item  `json:"items"`
	Fields        Field   `json:"field,omitempty"`
	Discounts     float64 `json:"discounts,omitempty"`
	Tax           float64 `json:"tax,omitempty"`
	Shipping      float64 `json:"shipping,omitempty"`
	AmountPaid    float64 `json:"amount_paid,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	Terms         string  `json:"terms,omitempty"`
}

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int64   `json:"quantity"`
	Unit_cost   float64 `json:"unit_cost"`
}

type Field struct {
	Tax       string  `json:"tax"`
	Discounts float64 `json:"discounts"`
	Shipping  float64 `json:"shipping"`
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
