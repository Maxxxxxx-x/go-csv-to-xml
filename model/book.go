package model

// Title,Author,Genre,Height,Publisher
type Book struct {
	Title     string `csv:"Title"       xml:"Title"`
	Author    string `csv:"Author"      xml:"Author"`
	Genre     string `csv:"Genre"       xml:"Genre"`
	Height    int    `csv:"Height"      xml:"Height"`
	Publisher string `csv:"Publisher"   xml:"Publisher"`
}

type BookList struct {
	Books []Book `xml:"Book"`
}
