# go-invoice-generator-connector
Go Connector to the Invoice-Generator API

###Example

```Go
r := new(Invoice)
r.From = "parag@invoiced.com"
r.To = "jared@invoiced.com"
r.Date = "January 2nd 2014"
r.Logo = "https://invoiced.com/img/header-logo.png"

items := make([]Item, 1)
item := Item{}
item.Name = "Marketing"
item.Quantity = 221
item.Unit_cost = 50.45
items[0] = item

r.Items = items

r.Create("bill", "/usr/local/pdfstash/")
```
