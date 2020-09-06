package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/pedroquerido/sword-challenge/notifications-agent/pkg/task"
	"github.com/streadway/amqp"
)

const (
	eventTaskCreated = "tasks.task.created"
)

type handler func(amqp.Delivery)

func handleCreatedTask(event amqp.Delivery) {

	task := &task.Task{}
	if err := json.Unmarshal(event.Body, task); err != nil {

		log.Printf("error unmarshalling task: %s\n", err.Error())
		if err := event.Reject(false); err != nil {
			log.Printf("error rejecting message: %s\n", event.MessageId)
		}
	}

	log.Printf("new task: %s performed task %s at %s", task.UserID, task.ID, task.Date.String())
	if err := event.Ack(false); err != nil {
		log.Printf("error acknoledging message: %s\n", event.MessageId)
	}
}
