package robot

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

	if s, ok := os.LookupEnv("QUANTITY_MESSAGE_QUEUE"); ok {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			v = 50
		}
		st.QuantityMessageQueue = uint16(v)
	}

	if s, ok := os.LookupEnv("RABBIT_MQ_URL"); ok {
		st.RabbitMqURL = s
	}

	return st
}
