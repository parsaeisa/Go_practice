# Pattern

In this documentation you can see useful golang patterns. Patterns such as:
- [Transaciton-generator](https://github.com/parsaeisa/Go_practice/blob/main/pattern/Transaction-Generator.md)
- [Opts](https://github.com/parsaeisa/Go_practice/blob/main/pattern/Opts.md)
- [Middleware](https://github.com/parsaeisa/Go_practice/blob/main/pattern/middleware.md)

## Other little patterns 

### Return cleaning methods

```go
func StartSampleConnection(... arguments) func() {
    // Strat the connection

   return func(){
        // Cleaning methods such as un-subscribes, Stop(), Clear(), etc.
   }
}
```