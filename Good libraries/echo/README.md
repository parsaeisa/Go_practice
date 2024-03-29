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

## CORS middleware

```go
s.e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
    AllowOrigins:     config.C.HTTP.CORS.AllowedOrigins,
    AllowMethods:     config.C.HTTP.CORS.AllowedMethods,
    AllowHeaders:     config.C.HTTP.CORS.AllowedHeaders,
    AllowCredentials: config.C.HTTP.CORS.AllowCredentials,
    ExposeHeaders:    config.C.HTTP.CORS.ExposedHeaders,
    MaxAge:           config.C.HTTP.CORS.MaxAge,
}))
```

## Recover with config middleware
```go
s.e.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
    StackSize:         config.C.HTTP.Recover.StackSize << 10,
    DisableStackAll:   config.C.HTTP.Recover.DisableStackAll,
    DisablePrintStack: config.C.HTTP.Recover.DisablePrintStack,
}))
```