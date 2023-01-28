# golang_otel

Here you can see my practice for open telemetry in golang . 
this is an example of nested span between server and client 
/ In this case we have grpc server and client . 

this program has some parts : 
+ setup 
+ monitoring 

## setup
in this stage you should define :
+ jeager agent 
+ trace provider


and we should register our provider with otel.SetTracerProvider() method . 
## monitoring 
For monitoring you should get a tracer from otel.GetTracerProvider().Tracer(app.Name)

It returns the tracer that you set in setup stage ( in app directory ) 

Then you define a span , with tracer.start() method . You give it a name and defer its cleaning method or use it any time that you are done with it . 

### install
for installing these methods libraries just import them , and then execute : 
``` bash 
go mod tidy
```
* don't use go get ... 


```go
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

var grpcClientCmd = &cobra.Command{
	Use: "client",
	Run: runClient,
}

func runClient(_ *cobra.Command, _ []string) {

	exp, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost("localhost"),
		jaeger.WithAgentPort("6831")))
	if err != nil {
		log.Logger.Fatal("error in setting jaeger endpoint", zap.Error(err))
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Logger.Fatal("error in getting hostname", zap.Error(err))
	}

	traceProvider := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// Record information about this application in a resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceInstanceIDKey.String(hostname),
			semconv.ServiceNameKey.String("client"),
			semconv.ServiceVersionKey.String("0.01"),
		)),
	)

	deferFunc := func() {
		// Do not make the application hang when it is shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := traceProvider.Shutdown(ctx)
		if err != nil {
			log.Logger.Fatal("error in shutting down telemetry", zap.Error(err))
		}
	}

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	fmt.Print("setup otel")
	defer deferFunc()

	c, conn := client.GetClient("localhost:50052")
	defer conn.Close()

	// ================================================
	ctxx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"user", "someone",
	))
	ctx, cancel := context.WithCancel(ctxx)
	defer cancel()

	tracer := otel.GetTracerProvider().Tracer("client")
	_, span := tracer.Start(ctx, "client_caller")
	defer span.End()

	time.Sleep(600 * time.Millisecond)
	// ================================================

	reply, err := c.SayHello(ctx, &kangaroo.Request{Name: "hello"})
	if err != nil {
		println(err.Error())
	}
	println(reply.Message)

	//app.Wait()
}

```

We have something like http middleware called interceptor in open telemetry . 
For each server and client we have two kinds of interceptor , 
* Unary
* Stream

So we have 4 methods : 
* UnaryClientInterceptor
* StreamClientInterceptor

* UnaryServerInterceptor
* StreamServerInterceptor

Its so simple to use otel in repositories . ( like redis and sqlx ) 
## Redis 
All you need to do is below line after defining redis client  : 
```go
a.RedisClient.AddHook(redisotel.NewTracingHook())
```

## Database 
And for database , while defining it you should use below method instead of sqlx : 
```go
db, err = otelsqlx.Open(config.C.Database.Driver, config.C.Database.DSN(), otelsql.WithAttributes(semconv.DBSystemMariaDB))
if err != nil {
	log.Fatalf("Cannot open database %s: %s", config.C.Database.String(), err)
}
```

These lines trace all of your commands on database and redis . 

## Http server echo
To wrap echo http server with otel , we need to define a middleware :
```python
s.e.Use(otelecho.Middleware(app.Name, otelecho.WithPropagators(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))))
```
`s.e` is a place where we define handlers , like `s.e.GET()` 
