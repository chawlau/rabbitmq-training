package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	} else {
		fmt.Println("send suc")
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.31.191:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := "Hello World! RabbitMQ Topic Exchange Message..."
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	exchangeName := "test_topic_exchange"
	routingKey1 := "user2.save"
	routingKey2 := "user2.update"
	routingKey3 := "user2.delete.abc"
	err = ch.Publish(exchangeName, routingKey1, false, false, msg)
	err = ch.Publish(exchangeName, routingKey2, false, false, msg)
	err = ch.Publish(exchangeName, routingKey3, false, false, msg)
	failOnError(err, "Failed to publish a message")
	time.Sleep(time.Duration(1) * time.Second)
	log.Printf(" [x] Sent %s", body)
}
