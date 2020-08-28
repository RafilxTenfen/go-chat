package app

import (
	"fmt"

	"github.com/leekchan/accounting"
)

var (
	ac = accounting.DefaultAccounting("$", 2)
)

// Stock represents a stock data, doesn't save in the database
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

// PublishFormat returns the stock formatted as “APPL.US quote is $93.42 per share”
func (s Stock) PublishFormat() string {
	return fmt.Sprintf("%s quote is %s per share", s.Symbol, s.ShareValue())
}

// ShareValue returns the share formated close value
func (s Stock) ShareValue() string {
	return ac.FormatMoneyFloat64(s.Close)
}
