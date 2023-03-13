package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/thanders/quiz-startup/broker/server/event"

	amqp "github.com/rabbitmq/amqp091-go"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Config is the type we'll use as a receiver to share application
// configuration around our app.
type Config struct {
	Rabbit         *amqp.Connection
	Etcd           *clientv3.Client
	LogServiceURLs map[string]string
}

// Payload is the type for data we push into RabbitMQ
type Payload struct {
	ServiceName string `json:"servicename"`
	Data        any    `json:"data"`
}

// connectToRabbit tries to connect to RabbitMQ, for up to 30 seconds
func connectToRabbit() (*amqp.Connection, error) {
	var rabbitConn *amqp.Connection
	var counts int64
	// var rabbitURL = os.Getenv("RABBIT_URL")

	for {
		connection, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("rabbitmq not ready...")
			counts++
		} else {
			fmt.Println()
			rabbitConn = connection
			break
		}

		if counts > 15 {
			fmt.Println(err)
			return nil, errors.New("cannot connect to rabbit")
		}
		fmt.Println("Backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
	fmt.Println("Connected to RabbitMQ!")
	return rabbitConn, nil
}

func (app *Config) HandleSubmission(action string, payload *Payload) {
	fmt.Println("THE PAYLOAD", payload.ServiceName)

	switch action {
	case "log":
		app.pushToQueue(payload.ServiceName, payload)
	default:
		fmt.Println("Default case in switch")
	}
}

// pushToQueue pushes a message into RabbitMQ
func (app *Config) pushToQueue(name string, msg *Payload) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		log.Println(err)
		return err
	}

	payload := Payload{
		ServiceName: name,
		Data:        msg,
	}

	j, _ := json.MarshalIndent(&payload, "", "    ")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}
