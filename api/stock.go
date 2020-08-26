package api

import (
	"fmt"
	"net/http"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/gocarina/gocsv"
)

// ErrStockNotFound returns a base error if a stock isn't found
var ErrStockNotFound = fmt.Errorf("Stock not Found")

// Stock returns a app.Stock structure based on a symbol, like.: "aapl.us"
func Stock(symbol string) (*app.Stock, error) {
	stocks := []*app.Stock{}

	resp, err := http.Get(fmt.Sprintf("https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv", symbol))
	if err != nil {
		return nil, err
	}

	if err := gocsv.Unmarshal(resp.Body, &stocks); err != nil {
		return nil, ErrStockNotFound
	}

	if len(stocks) < 1 {
		return nil, ErrStockNotFound
	}

	return stocks[0], nil
}
