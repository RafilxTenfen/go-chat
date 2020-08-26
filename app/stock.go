package app

// Stock represents a stock data
type Stock struct {
	Symbol string
	Date   string
	Time   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume uint64
}
