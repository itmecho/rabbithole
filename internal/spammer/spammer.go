package spammer

import (
	"math/rand"
	"time"

	"github.com/itmecho/rabbithole/pkg/rabbitmqclient"
	log "github.com/sirupsen/logrus"
)

// SendConfig Stores the configuration for the spammer's Send function
type SendConfig struct {
	MinSize int
	MaxSize int
	Count   int
}

// Send Sends messages to the client queue based on the provided configuration
func Send(client rabbitmqclient.Client, config SendConfig) error {
	rand.Seed(time.Now().UnixNano())

	var count int

	for count < config.Count {
		var size int
		if config.MaxSize == config.MinSize {
			size = config.MinSize
		} else {
			size = rand.Intn(config.MaxSize-config.MinSize) + config.MinSize
		}

		body := make([]byte, size)

		log.Debugf("Generating message of %d bytes", size)

		_, err := rand.Read(body)
		if err != nil {
			return err
		}

		log.Debugf("sending message #%d", count)
		if err = client.Publish("text/plain", body); err != nil {
			return err
		}
		count++
	}
	return nil
}
