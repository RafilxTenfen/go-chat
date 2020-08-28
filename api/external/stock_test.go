package external_test

import (
	"strings"
	"testing"

	"github.com/RafilxTenfen/go-chat/api/external"
)

func TestStock(t *testing.T) {
	tests := []struct {
		symbol        string
		expectedError error
	}{
		{
			symbol:        "aapl.us",
			expectedError: nil,
		},
		{
			symbol:        "aadsaplx.us",
			expectedError: external.ErrStockNotFound,
		},
	}

	for i := range tests {
		test := tests[i]

		stock, err := external.Stock(test.symbol)
		if err != test.expectedError {
			t.Errorf("Error on test %d \nReceived Error:%s", i, err.Error())
		}

		if stock == nil {
			continue
		}

		if stock.Symbol != strings.ToUpper(test.symbol) {
			t.Errorf("Error on test %d \nExpected Stock:%s \nReceived Stock:%+v", i, test.symbol, stock.Symbol)
		}

	}
}
