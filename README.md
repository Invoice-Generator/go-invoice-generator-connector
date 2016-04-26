# go-invoice-generator-connector
Go Connector to the Invoice-Generator API

It lets you create a pdf invoice by querying the invoice-generator.com api.


###Example

```Go
    r := new(Invoice)

    r.From = "parag@invoiced.com"
    r.To = "jared@invoiced.com"
    r.Date = "January 2nd 2014"
    r.Logo = "https://invoiced.com/img/header-logo.png"
    r.PaymentTerms = "NET 30"
    r.Number = "INV-001"
    r.PurchaseOrder = "PO-32422"

    items := make([]Item, 1)
    item := Item{}
    item.Name = "Marketing"
    item.Quantity = 221
    item.Unit_cost = 50.45

    items[0] = item

    r.Items = items

    r.Discounts = 12.21
    r.Tax = 13.11
    r.Shipping = 12.20

    r.TaxTitle = "Austin Tax"

    r.Currency = "UAH"

    err := r.Create("/usr/local/pdfstash/bill7.pdf")

    if err != nil {
        panic(err)
    }
```
