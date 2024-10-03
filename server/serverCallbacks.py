import bgRemove

def onRequestReceived(ch, method, properties, body) :
    print(f"Request received, corr_id: {properties.correlation_id}, reply_to: {properties.reply_to}")
    value = bgRemove.removeRequest(body)
    ch.basic_publish("", routing_key=properties.reply_to, body=value)