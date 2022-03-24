package models

type StockData struct {
	Date      string  `csv:"Date"      json:"Date"      bson:"Date"`
	Open      float64 `csv:"Open"      json:"Open"      bson:"Open"`
	High      float64 `csv:"High"      json:"High"      bson:"High"`
	Low       float64 `csv:"Low"       json:"Low"       bson:"Low"`
	Close     float64 `csv:"Close"     json:"Close"     bson:"Close"`
	Adj_Close float64 `csv:"Adj Close" json:"Adj Close" bson:"Adj Close"`
	Volume    int     `csv:"Volume"    json:"Volume"    bson:"Volume"`
}
