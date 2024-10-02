package messaging

import (
	"log"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/polatefekaya/photoBooth/internal/photo"
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
		  log.Printf("Received a message, %s", d.Body[:50])
			savePhoto(&d)
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
	

	id := uuid.New()


	pht := photo.Photo{
		Path: "./resources/image.jpg",
	}

	img, _, err := pht.DecodePhoto() 
	failOnError(err, "Failed to decode the photo")

	body, err := pht.EncodePhoto(&img)
	failOnError(err, "Failed to encode the photo")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        body,
		  ReplyTo: replyQueue.Name,
		  CorrelationId: id.String(),
	})
	failOnError(err, "Failed to publish a message")
	
	log.Printf(" [x] Sent %s\n", body[:20])

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
		  log.Printf("Received a message")
		  savePhoto(&d)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}

func savePhoto(d *amqp.Delivery){
	pht := photo.Photo{
		Path: "./resources/gen/image.jpeg",
	}

	pht.SavePhoto(&d.Body)
}