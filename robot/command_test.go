package robot_test

import (
	"testing"

	"github.com/RafilxTenfen/go-chat/robot"
	"github.com/streadway/amqp"
)

func TestIsCommand(t *testing.T) {
	tests := []struct {
		msg      string
		expected bool
	}{
		{
			msg:      "aapl.us",
			expected: false,
		},
		{
			msg:      "Received any message",
			expected: false,
		},
		{
			msg:      "/stock=stock_code",
			expected: true,
		},
		{
			msg:      "/othercommand=anyvalue",
			expected: true,
		},
	}

	for i, test := range tests {
		d := amqp.Delivery{
			Body: []byte(test.msg),
		}

		received := robot.IsCommand(d)
		if received != test.expected {
			t.Errorf("Error on test %d \nReceived: %t, Expected: %t", i, received, test.expected)
		}

	}
}

func TestGetCommand(t *testing.T) {
	tests := []struct {
		msg      string
		expected string
	}{
		{
			msg:      "/stock=stock_code",
			expected: "stock",
		},
		{
			msg:      "/othercommand=anyvalue",
			expected: "othercommand",
		},
		{
			msg:      "anyvaluemsg",
			expected: "anyvaluemsg",
		},
	}

	for i, test := range tests {
		d := amqp.Delivery{
			Body: []byte(test.msg),
		}

		received := robot.GetCommand(d)
		if received != test.expected {
			t.Errorf("Error on test %d \nReceived: %s, Expected: %s", i, received, test.expected)
		}

	}
}

func TestGetCommandValue(t *testing.T) {
	tests := []struct {
		msg      string
		expected string
	}{
		{
			msg:      "/stock=stock_code",
			expected: "stock_code",
		},
		{
			msg:      "/othercommand=anyvalue",
			expected: "anyvalue",
		},
		{
			msg:      "anyvaluemsg",
			expected: "anyvaluemsg",
		},
	}

	for i, test := range tests {
		d := amqp.Delivery{
			Body: []byte(test.msg),
		}

		received := robot.GetCommandValue(d)
		if received != test.expected {
			t.Errorf("Error on test %d \nReceived: %s, Expected: %s", i, received, test.expected)
		}

	}
}
