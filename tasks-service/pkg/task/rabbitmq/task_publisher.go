package rabbitmq

import (
	"encoding/json"

	"github.com/google/uuid"

	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"github.com/streadway/amqp"
)

var _ task.Publisher = (*taskPublisher)(nil)

type taskPublisher struct {
	ch                    *amqp.Channel
	exchange              string
	taskCreatedRoutingKey string
}

// NewTaskPublisher ...
func NewTaskPublisher(conn *amqp.Connection, exchangeName string, taskCreatedRoutingKey string) (task.Publisher, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := declareExchange(ch, exchangeName); err != nil {
		return nil, err
	}

	return &taskPublisher{
		ch:                    ch,
		exchange:              exchangeName,
		taskCreatedRoutingKey: taskCreatedRoutingKey,
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

func (p *taskPublisher) PublishTaskCreated(t *task.Task) error {

	body, err := json.Marshal(t)
	if err != nil {
		return pkgError.NewError(task.ErrorPublishingTaskEvent).WithDetails(err.Error())
	}

	err = p.ch.Publish(
		p.exchange,              // exchange
		p.taskCreatedRoutingKey, // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			MessageId:   uuid.New().String(),
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return pkgError.NewError(task.ErrorPublishingTaskEvent).WithDetails(err.Error())
	}

	return nil
}
