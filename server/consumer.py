import pika
import uuid

connParameters = pika.ConnectionParameters("localhost")
conn = pika.BlockingConnection(connParameters)

channel = conn.channel()

replyQueue = channel.queue_declare(queue="", exclusive=True)

def onCallBack(ch, method, properties, body):
    print(f"Message {body}")

channel.basic_consume(queue=replyQueue.method.queue, auto_ack=True, on_message_callback=onCallBack)

channel.queue_declare("request-queue")


corr_id = str(uuid.uuid4())
print(f"Sending Request: {corr_id}")

message = "naber knk"

channel.basic_publish(
    "",
    routing_key="request-queue",
    body=message,
    properties= 
        pika.BasicProperties(
            reply_to=replyQueue.method.queue,
            correlation_id=corr_id
        )
)

print(f"Starting Client")

channel.start_consuming()
