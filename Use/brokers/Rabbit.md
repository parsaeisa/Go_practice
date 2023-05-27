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

## Exchanges
Creating exchanges with golang:
```go
exchangeName := "myExchange"
exchangeType := "direct"

err = ch.ExchangeDeclare(
    exchangeName, // Name of the exchange
    exchangeType, // Type of the exchange (e.g., direct, fanout, topic)
    false,        // Durable - exchange survives broker restart
    false,        // Auto-deleted - exchange is deleted when there are no bindings
    false,        // Internal - only used for exchange-to-exchange bindings
    false,        // No-wait - don't wait for a response
    nil,          // Arguments
)
if err != nil {
    log.Fatal(err)
}
```

## Consuming

For consuming in golang: 
```go
msgs, err := ch.Consume(
    queueName, // Queue name to consume from
    "",        // Consumer tag - unique identifier for the consumer
    true,      // Auto-acknowledge - messages are automatically acknowledged
    false,     // Exclusive - this consumer gets exclusive access to the queue
    false,     // No-local - don't deliver messages published by this connection
    false,     // No-wait - don't wait for a response
    nil,       // Arguments
)
if err != nil {
    log.Fatal(err)
}

// Start a goroutine to handle incoming messages
go func() {
    for msg := range msgs {
        // Process the received message
        log.Printf("Received message: %s", msg.Body)
    }
}()

// Block the main goroutine to keep the consumer alive
select {}

```

As you can see we should create a new goroutine to consume on a queue.

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

