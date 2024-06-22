package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"TweetDelivery/config"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var rabbitmqConfig = config.NewRabbitMQConfig()

func CreateQueue(user string) {
	// Connect to RabbitMQ Server
	conn, err := amqp.Dial(rabbitmqConfig.GetURL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	// TODO: Ideally a retry logic on error should be implemented
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		rabbitmqConfig.GetExchangeName(),
		rabbitmqConfig.GetExchangeType(),
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to create exchange")

	_, err = ch.QueueDeclare(
		user,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue for")
}

func CreateBinding(user string, follower string) error {
	// Connect to RabbitMQ Server
	conn, err := amqp.Dial(rabbitmqConfig.GetURL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	// TODO: Ideally a retry logic on error should be implemented
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	log.Printf("Binding queue %s to exchange %s with routing key %s", user, rabbitmqConfig.GetExchangeName(), follower)
	return ch.QueueBind(
		user,
		follower,
		rabbitmqConfig.GetExchangeName(),
		false,
		nil,
	)
}

func PublishTweet(user string, tweet string) error {
	// Connect to RabbitMQ Server
	conn, err := amqp.Dial(rabbitmqConfig.GetURL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	// TODO: Ideally a retry logic on error should be implemented
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(ctx,
		rabbitmqConfig.GetExchangeName(),
		user,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(tweet),
		},
	)
}

func ConsumeTweet(user string) []string {
	timeout := 5 * time.Second
	// Connect to RabbitMQ Server
	conn, err := amqp.Dial(rabbitmqConfig.GetURL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	// TODO: Ideally a retry logic on error should be implemented
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		user,
		"",
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Error consuming tweets: %v", err)
	}

	var tweetsHome []string
	timeoutCh := time.After(timeout)

	for {
		select {
		case msg := <-msgs:
			fmt.Println(string(msg.Body))
			tweetsHome = append(tweetsHome, string(msg.Body))
		case <-timeoutCh:
			fmt.Println("Timeout reached, stopping consumption")
			return tweetsHome
		}
	}
}
