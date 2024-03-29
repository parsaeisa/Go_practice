# encoding

## encoding/base64
```go
secureKey := base64.URLEncoding.EncodeToString(token)
decToken, err = base64.URLEncoding.DecodeString(apiKey)
```

## encoding/binary

```go
binary.BigEndian.Uint16(b[1:])
```