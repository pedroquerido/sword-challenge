package rabbitmq

import (
	"log"

	"github.com/pedroquerido/sword-challenge/notifications-agent/internal/listener"
	"github.com/streadway/amqp"
)

type eventListener struct {
	ch       *amqp.Channel
	queue    amqp.Queue
	handlers map[string]handler
}

// NewEventListener ...
func NewEventListener(conn *amqp.Connection, exchangeName, queueName string, bindingKeys ...string) (listener.EventListener, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := declareExchange(ch, exchangeName); err != nil {
		return nil, err
	}

	queue, err := setupQueue(ch, queueName)
	if err != nil {
		return nil, err
	}

	handlers, err := declareBindings(ch, exchangeName, queueName, bindingKeys...)
	if err != nil {
		return nil, err
	}

	return &eventListener{
		ch:       ch,
		queue:    queue,
		handlers: handlers,
	}, nil
}

func declareExchange(ch *amqp.Channel, exchangeName string) error {

	return ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func setupQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func declareBindings(ch *amqp.Channel, exchangeName, queueName string, bindingKeys ...string) (map[string]handler, error) {

	handlers := make(map[string]handler, len(bindingKeys))
	for _, key := range bindingKeys {

		// bind key
		err := ch.QueueBind(
			queueName,    // queue name
			key,          // routing key
			exchangeName, // exchange
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			return nil, err
		}

		// register handler
		switch key {
		case eventTaskCreated:
			handlers[eventTaskCreated] = handleCreatedTask
		}
	}

	return handlers, nil
}

func (l *eventListener) Listen() error {

	events, err := l.ch.Consume(
		l.queue.Name, // queue
		"",           // consumer
		false,        // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
	if err != nil {
		return err
	}

	go func() {
		for event := range events {

			handler, ok := l.handlers[event.RoutingKey]
			if !ok {
				if err := event.Reject(false); err != nil {
					log.Printf("error rejecting message: %s\n", event.MessageId)
				}
			}

			handler(event)
		}
	}()

	return nil
}
