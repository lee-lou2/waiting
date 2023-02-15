package que

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func New() (*amqp.Channel, error) {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.Channel()
}

func Message(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"qr",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func Publish(ch *amqp.Channel, message string) error {
	q, err := ch.QueueDeclare(
		"qr",  // queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
