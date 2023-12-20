# Graceful Shutdown

The main idea of the Graceful Shutdown pattern is shutting anything down takes time and we should wait for that time. 

A channel is needed to inform application that it is about to being shutted down. Call this channel `quite`.

For example for shutting down an http server, there should be a channel where you announce your **decistion to shutdown the application** in, and another channel for announcing that **the process of shutting down is completed**. 

As you know, for waiting on a channel to get a message in you can use '<-[channel name]'. 

```go

quit := make(chan struct{})

// first channel
httpShutdownRequest := make(chan struct{})

// second channel
httpShutdownReady := make(chan struct{})

go func(){
    // starting your http server
    // ...
}

go func(){
    // waiting for signals
    <-httpShutdownRequest    
    // start shutting down the http server

    // announce that shutting down is completed
    httpShutdownReady <- struct{}{}
}

// The process of shutting down the application

// wait on channel to get a message for shutting down
<- quit

// announce http shutdown
httpShutdownRequest <- struct{}{}

// waiting until the shutdown process is completed
<- httpShutdownReady

// shutting down other dependencies

```


## Supported graceful shutdowns

Many packages have built-in methods for graceful shutdown and you don't have to implement it manually. 

For brokers such as Nats, rabbit mq, etc. there methods that are like `Unsubscribe`, `UnsubscribeAll`, etc. which do the graceful shutdown. 

All `Close` for connections are being used for graceful shutdown. Most of the times, Close method is being called after `defer` for cleaning.

