package rabbit

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rhizomplatform/log"
)

// Settings define basic settings for initiate a Bot
type Settings struct {
	QuantityMessageQueue uint16
	RabbitMqURL          string
}

// LoadSettingsFromEnv returns a Settings based on .env
func LoadSettingsFromEnv() Settings {
	_ = godotenv.Load(".env")

	var st Settings

	st.QuantityMessageQueue = getQuantityMessageQueue()

	if s, ok := os.LookupEnv("RABBIT_MQ_URL"); ok {
		st.RabbitMqURL = s
	}

	return st
}

func getQuantityMessageQueue() uint16 {
	if s, ok := os.LookupEnv("QUANTITY_MESSAGE_QUEUE"); ok {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			// default value
			log.WithError(err).Error("error on convert QUANTITY_MESSAGE_QUEUE")
			return 50
		}
		return uint16(v)
	}
	return 50

}
