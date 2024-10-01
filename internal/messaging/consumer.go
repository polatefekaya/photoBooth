package messaging

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}
//amqp://guest:guest@localhost:5672
func StartConnection(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	replyQueue, err := ch.QueueDeclare(
		"", // name
		false,   // durable
		false,   // delete when unused
		true,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	msgs, err := ch.Consume(
		replyQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )
	  failOnError(err, "Failed to register a consumer")
	  

	  go func() {
		for d := range msgs {
		  log.Printf("Received a message: %s", d.Body)
		}
	  }()
	  
	  


	q, err := ch.QueueDeclare(
		"request-queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		  ReplyTo: replyQueue.Name,
		  CorrelationId: "1",
	})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
	msgs2, err := ch.Consume(
		replyQueue.Name, // queue
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
		for d := range msgs2 {
		  log.Printf("Received a message: %s", d.Body)
		  
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}