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

	exchangeName := "test_direct_exchange"
	exchanegType := "direct"
	queueName := "test_direct_queue"
	routingKey := "test.direct"

	//声明交换机
	err = ch.ExchangeDeclare(exchangeName, exchanegType, false, false, false, false, nil)
	failOnError(err, "Failed to declare a exchange")
	//声明队列
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	//绑定
	err = ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	failOnError(err, "Failed to bind a queue")

	forever := make(chan bool)

	//consumerTag := "direct_consumer"
	go func() {
		for {
			delivery, err := ch.Consume(
				queueName,  // queue
				"consumer", // consumer
				true,       // auto-ack
				false,      // exclusive
				false,      // no-local
				false,      // no-wait
				nil,        // args
			)
			failOnError(err, "Failed to register a consumer")

			for d := range delivery {
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
