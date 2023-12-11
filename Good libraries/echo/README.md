# Echo

Echo has custom validator. 

## Proxy middleware

When defining a route in echo, you can add a middleware that redirects it's requests between multiple targets. While sending requests to multiple targets, it can **balance the load** with different **load balancing algorithms** such as round robin. 

This is an example:
```go
    s.e.Group("/route ", middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
            {
                URL: url1,
            },
        })))
```

As you can see, you can not add a handler.