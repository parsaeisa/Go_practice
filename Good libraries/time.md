# Time

Golang built-in package.

Good methods: 

- `time.Since()`

## Change timezone

```go
var timezone time.Location
timezone = time.UTC

time, err = time.ParseInLocation(TimeLayout, rawTime, timezone)
if err != nil {
    // TODO: Shouldn't we return error here ?
    log.WithError(err).Error("error in parsing accepted event time")
}
```