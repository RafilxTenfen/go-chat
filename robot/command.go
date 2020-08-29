package robot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RafilxTenfen/go-chat/api/external"
	"github.com/RafilxTenfen/go-chat/app"
	"github.com/streadway/amqp"
)

var (
	regAllCommandMsg = regexp.MustCompile(`^\\/\\w+=[A-Za-z0-9_.]+`)
	regCommand       = regexp.MustCompile(`\\w+`)
	regCommandValue  = regexp.MustCompile(`[A-Za-z0-9_.]+$`)
)

// HandleReceivedCommand receives a command
func (b *Bot) HandleReceivedCommand(d amqp.Delivery) error {
	command := string(regCommand.Find(d.Body))

	switch strings.ToLower(command) {
	case "stock":
		return b.Stock(d)
	default:
		return fmt.Errorf("Command %s not handle", command)
	}
}

// Stock Bot Command searches for stock data and print in the given queue
func (b *Bot) Stock(d amqp.Delivery) error {
	symbol := GetCommandValue(d)

	stock, err := external.Stock(symbol)
	if err != nil {
		return err
	}

	return b.PublishStock(d.RoutingKey, stock)
}

// PublishStock publish a message in the queue based on the stock
func (b *Bot) PublishStock(queueName string, stock *app.Stock) error {
	queue, ok := b.queueMap.Load(queueName)
	if !ok {
		return app.ErrQueueNotFound
	}

	return queue.Publish(b.channel, stock.PublishFormat())
}

// IsCommand returns true if the command matches with something like "/stock=stock_code"
func IsCommand(d amqp.Delivery) bool {
	return regAllCommandMsg.Match(d.Body)
}

// GetCommand returns the command as string, receiving "/stock=stock_code" returns "stock"
func GetCommand(d amqp.Delivery) string {
	return string(regCommand.Find(d.Body))
}

// GetCommandValue returns the command value as string, receiving "/stock=aapl.us" returns "aapl.us"
func GetCommandValue(d amqp.Delivery) string {
	return string(regCommandValue.Find(d.Body))
}
