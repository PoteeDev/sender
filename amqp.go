package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PoteeDev/sender/providers"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitClient struct {
	Login    string
	Password string
	Address  string
	Port     int
}

func NewRabbitClient() *RabbitClient {
	return &RabbitClient{
		Login:    os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASS"),
		Address:  os.Getenv("RABBITMQ_HOST"),
		Port:     5672,
	}
}

func (rc *RabbitClient) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", rc.Login, rc.Password, rc.Address, rc.Port)
}

func AMQPReciver() {
	conn, err := amqp.Dial(NewRabbitClient().URL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sender", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var message JsonMessage
			// decode json data
			err := json.NewDecoder(bytes.NewBuffer(d.Body)).Decode(&message)
			if err != nil {
				log.Println(err)
			}
			log.Printf("received a message: %s", message)
			token := os.Getenv("BOT_TOKEN")
			err = providers.NewMessage(token, message.ChatID, message.Message, "").Send()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
