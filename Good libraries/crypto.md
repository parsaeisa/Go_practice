# Crypto

Golang has a built-in package called crypto. 

Here we see some examples of this package. 

## crypto/cipher

```go
// I think AEAD is a encryption algorithm. Like AES.
var cipher cipher.AEAD
encToken := cipher.Seal(nonce, nonce, a.sign(token.Pack()), nil)
secureKey := base64.URLEncoding.EncodeToString(encToken)
```