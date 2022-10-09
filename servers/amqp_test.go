package servers

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestSendAmqp(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testData := map[string]string{
		"recipient": "naliway",
		"message":   "hello, amqp!",
	}
	b, _ := json.Marshal(testData)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", testData)
}
