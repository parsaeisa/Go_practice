# Rabbit

## getting connection

```go
import (
    "log"
    "time"
    "github.com/streadway/amqp"
)

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

## Creating channel

We can create channel from a connection :
```go
ch, err := conn.Channel()
if err != nil {
    log.Fatal(err)
}
defer ch.Close()
```

## Adding a delayed job

For adding a job to Rabbit, we need an exchange and a Queue, then we need to bind them.

```go
delayedQueueName := "delayedQueue"
delayedExchangeName := "delayedExchange"

// Declare the delayed queue
_, err = ch.QueueDeclare(
    delayedQueueName, // Name of the delayed queue
    true,             // Durable - message survives broker restart
    false,            // Auto-delete - queue is deleted when there are no consumers
    false,            // Exclusive - queue can be accessed by this connection only
    false,            // No-wait - don't wait for a response
    amqp.Table{
        "x-dead-letter-exchange":    "",                // DLX for expired messages
        "x-dead-letter-routing-key": "originalRoutingKey", // Routing key for expired messages
    }, // Arguments
)
if err != nil {
    log.Fatal(err)
}

// Declare the delayed exchange
err = ch.ExchangeDeclare(
    delayedExchangeName, // Name of the delayed exchange
    "direct",            // Type of the exchange (direct)
    true,                // Durable - exchange survives broker restart
    false,               // Auto-deleted - exchange is deleted when there are no bindings
    false,               // Internal - only used for exchange-to-exchange bindings
    false,               // No-wait - don't wait for a response
    nil,                 // Arguments
)
if err != nil {
    log.Fatal(err)
}

// Bind the delayed queue to the delayed exchange
err = ch.QueueBind(
    delayedQueueName,     // Name of the delayed queue
    "originalRoutingKey", // Routing key for the delayed queue
    delayedExchangeName,  // Name of the delayed exchange
    false,                // No-wait - don't wait for a response
    nil,                  // Arguments
)
if err != nil {
    log.Fatal(err)
}
```

