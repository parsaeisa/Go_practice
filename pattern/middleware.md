# Middleware pattern

I think there is no such thing called middleware, it's just a pattern that I like to remember. 

This pattern is kind of functional programming. 

```go
func XMiddleware(next handlerFunc) handlerFunc {
    return func(ctx context.Context) error {
        if /* any guard detection */ {
            return next()
        }

        // pre-processing 

        err := next(ctx)
        if err != nil {
            // ....
        }

        // post-processing 
    }
}
```