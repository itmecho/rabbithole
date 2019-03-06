package rabbitmqclient

import "github.com/streadway/amqp"

// Client Stores connection and information for rabbitmq
type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewClient Creates a new Client storing the connection, channel, and queue
func NewClient(url string, queue string) (client Client, err error) {
	client.conn, err = amqp.Dial(url)
	if err != nil {
		return
	}

	client.channel, err = client.conn.Channel()
	if err != nil {
		return
	}

	client.queue, err = client.channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	return
}

// Publish Publishes a message to the client's queue
func (cli Client) Publish(contentType string, body []byte) error {
	return cli.channel.Publish("", cli.queue.Name, false, false, amqp.Publishing{
		ContentType: contentType,
		Body:        body,
	})
}

// Close Shuts down the client's channel and connection
func (cli Client) Close() {
	cli.channel.Close()
	cli.conn.Close()
}
