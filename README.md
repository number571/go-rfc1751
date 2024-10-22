# go-rfc1751

> https://datatracker.ietf.org/doc/html/rfc1751

The RFC1751 encoding in Go language. The code is mostly rewritten from the Python [implementation](https://github.com/remram44/python-rfc1751). Tested on examples of repositories [remram44/python-rfc1751](https://github.com/remram44/python-rfc1751), [vmizg/rfc1751.js](https://github.com/vmizg/rfc1751.js), [arokettu/php-rfc1751](https://github.com/arokettu/php-rfc1751).

### Example

```go
b := []byte{204, 172, 42, 237, 89, 16, 86, 190}
fmt.Println(EncodeToString(b))
// Output: RASH BUSH MILK LOOK BAD BRIM
```

```go
s := "RASH BUSH MILK LOOK BAD BRIM"
fmt.Println(DecodeString(s))
// Output: [204 172 42 237 89 16 86 190] <nil>
```
