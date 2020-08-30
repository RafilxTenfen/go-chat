package rabbit

import (
	"fmt"
	"time"

	"github.com/rhizomplatform/log"
	"github.com/streadway/amqp"
)

// Init starts a connection with rabbitMQ
func Init(url string) (*amqp.Connection, *amqp.Channel, error) {
	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		log.With(log.F{
			"tries": i + 1,
			"url":   url,
		}).Info("Connecting to rabbitMQ")
		conn, err := amqp.Dial(url)
		if err != nil {
			log.Error(err)
			continue
		}
		ch, err := conn.Channel()
		if err != nil {
			return nil, nil, err
		}

		return conn, ch, nil
	}

	return nil, nil, fmt.Errorf("Error on connect to rabbitMQ")
}
