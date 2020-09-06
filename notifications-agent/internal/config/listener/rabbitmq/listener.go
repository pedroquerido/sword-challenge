package rabbitmq

import (
	"github.com/pedroquerido/sword-challenge/notifications-agent/internal/listener"
	"github.com/streadway/amqp"
)

type eventListener struct {
	conn  *amqp.Connection
	queue string
}

// NewListener ...
func NewListener(conn *amqp.Connection, queue, exchange string, bindingKeys ...string) listener.EventListener {

	l := &eventListener{
		conn: conn,
	}

}

func (l *eventListener) Listen() error {

	return nil
}
