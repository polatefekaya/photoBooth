package messaging

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/polatefekaya/photoBooth/internal/photo"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func warnOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

//amqp://guest:guest@localhost:5672

func openChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	warnOnError(err, "Issue while opening a channel")
	return ch, err
}

func declareReplyQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	replyQueue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	warnOnError(err, "Issue while declaring a reply queue")
	return &replyQueue, err
}

func consumeReplyQueue(ch *amqp.Channel, qname string) {
	msgs, err := ch.Consume(
		qname, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message, %s", d.Body[:50])
			savePhoto(&d)
		}
	}()
}

func declareRequestQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"request-queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	warnOnError(err, "Issue while declaring a request queue")
	return &q, err
}

func photoBytes() []byte {
	pht := photo.Photo{
		Path: "./resources/image.jpg",
	}

	img, _, err := pht.DecodePhoto()
	failOnError(err, "Failed to decode the photo")

	body, err := pht.EncodePhoto(&img)
	failOnError(err, "Failed to encode the photo")

	return body
}

func publish(ch *amqp.Channel, reqQueueName string, replyQueueName string) error {
	id := uuid.New()
	fmt.Printf("new uuid is: %s\n", id.String())

	body := photoBytes()

	err := ch.Publish(
		"",           // exchange
		reqQueueName, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          body,
			ReplyTo:       replyQueueName,
			CorrelationId: id.String(),
		})
	warnOnError(err, "Issue while publishing")
	return err
}

func StartConnection() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := openChannel(conn)
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	replyQueue, err := declareReplyQueue(ch)
	fmt.Printf("Declared Reply Queue Name: %s\n", replyQueue.Name)
	failOnError(err, "Failed to declare a queue")

	consumeReplyQueue(ch, replyQueue.Name)

	reqQueue, err := declareRequestQueue(ch)
	failOnError(err, "Failed to declare a queue")

	err = publish(ch, reqQueue.Name, replyQueue.Name)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent an image\n")

	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func savePhoto(d *amqp.Delivery) {
	pht := photo.Photo{
		Path: "./resources/gen/image.png",
	}

	pht.SavePng(&d.Body)
}
