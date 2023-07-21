package main

import (
	"fmt"
	"log"
	"strings"

	//jiyas package
	"github.com/golang-module/dongle"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"EMAILQUEUE",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetchCount
		0,     // prefetchSize
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			data := string(d.Body)
			decy_data := dongle.Decode.FromString(data).ByBase64().ToString() // hello world

			//decrypt

			parts := strings.Split(decy_data, ";")

			if len(parts) != 2 {
				fmt.Println("Invalid data format")
				return
			}

			msg_type, email_add := parts[0], parts[1]
			//split data into phone and message type
			fmt.Println("data1 =", msg_type)
			fmt.Println("data2 =", email_add)
			//fathersms{ msgtype ; email_add} { !! pending}
			FatherEmail(msg_type, email_add)
			log.Printf("Done")

			d.Ack(false) // Acknowledge the message
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	<-make(chan struct{})
}
