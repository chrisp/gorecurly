package gorecurly

type Address struct {
    Address1 string `xml:"address1,omitempty"`
    City     string `xml:"city,omitempty"`
    Country  string `xml:"country,omitempty"`
    Zip      string `xml:"zip,omitempty"`
}
