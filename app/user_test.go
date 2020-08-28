package app_test

import (
	"testing"

	"github.com/RafilxTenfen/go-chat/app"
)

func TestValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{
			email:    "anyea.dsasa",
			expected: false,
		},
		{
			email:    "táo@hotmail.com",
			expected: false,
		},
		{
			email:    "rafilx.dsasa@gmail.com",
			expected: true,
		},
		{
			email:    "tao.dsasa@hotmail.com",
			expected: true,
		},
	}

	for i, test := range tests {
		received := app.ValidEmail(test.email)
		if test.expected != received {
			t.Errorf("Error on test %d \nExpected: %t \nReceived: %t", i, test.expected, received)
		}
	}
}
