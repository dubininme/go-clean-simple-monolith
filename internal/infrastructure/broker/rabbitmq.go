package broker

import (
	"fmt"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewRabbitMQ(user, password, host string, port int) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (r *RabbitMQ) Close() error {
	if r.channel != nil {
		err := r.channel.Close()
		if err != nil {
			return err
		}
	}
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}
