package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // autoDelete when done?
		false,        // only internal exchange?
		false,        // noWait
		nil,          //args
	)
}

func declareRandomQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		"",    // name
		false, // durable - delete when done
		false, // autoDelete - delete when unused
		true,  // exclusive - dont use this around
		false, // noWait
		nil,   // args
	)
}
