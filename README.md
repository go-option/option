# Golang function options helper

## `import "gopkg.in/option.v0"`

![Go Import](https://img.shields.io/badge/import-gopkg.in/option.v0-9cf?logo=go&style=for-the-badge)
[![Go Reference](https://img.shields.io/badge/reference-go.dev-007d9c?logo=go&style=for-the-badge)](https://pkg.go.dev/gopkg.in/option.v0)

## Example

```go
import "gopkg.in/option.v0"

// Define an opaque option type like this
type CreationOption func(*creationOptions)

// Define the concrete option properties as a struct type
type creationOptions struct {
	conn *net.Conn
}

// Define functions that generate closures that sets the
// option properties on the struct type
func CreateWithConn(conn *net.Conn) CreationOption {
	return func(o *creationOptions) {
		o.conn = conn
	}
}

// Now to use it in actual API functions, add it as the last
// variadic parameter.
func Create(options ...CreationOption) *Client {
    // Use the helper function from this module
    opts := option.New(options)
    // Use the returned options struct value with all user
    // provided values
    return &Client{opts.conn}
}
```
