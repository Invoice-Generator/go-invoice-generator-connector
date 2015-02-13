package invoicegenerator

import (
	"testing"
)

func TestCreate(t *testing.T) {

	r := new(Invoice)
	r.From = "parag@invoiced.com"
	r.To = "jared@invoiced.com"
	r.Create()
}
