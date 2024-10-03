import pika
import serverCallbacks as callbacks

class server:
    connParameters = pika.ConnectionParameters("localhost")
    conn = pika.BlockingConnection(connParameters)

    channel = conn.channel()

    channel.queue_declare("request-queue")
    
    channel.basic_consume("request-queue", auto_ack=True, on_message_callback=callbacks.onRequestReceived)

    print("Started to consume")

    channel.start_consuming()
