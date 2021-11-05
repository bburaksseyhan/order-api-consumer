package pkg

import (
	"github.com/bburaksseyhan/orderconsumer/src/cmd/utils"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

const (
	QUEUENAME = "NEW_ORDER_QUEUE"
)

func Initialize(config utils.Configuration) {

	log.Info("Consumer is running!!!", config.Queue.Url)

	conn, err := amqp.Dial(config.Queue.Url)
	if err != nil {
		log.Info("Failed Initializing Broker Connection", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Info(err)
	}
	defer ch.Close()

	messages, err := ch.Consume(QUEUENAME, "", true, false, false, false, nil)

	if err != nil {
		log.Error(err)
	}

	channel := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("Recieved Order Message: %s\n", message.Body)
		}
	}()

	log.Info(" [*] - Waiting for messages")

	<-channel
}
