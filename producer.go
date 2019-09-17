package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.31.191:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	routingKey := "hello"
	body := "Hello World!"
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	for {
		err = ch.Publish(
			"",         // exchange
			routingKey, // routing key
			false,      // mandatory
			false,      // immediate
			msg,
		)
		//time.Sleep(time.Duration(1) * time.Second)
	}
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
