package common

import "time"

const (
	ReadTimeout = 120 * time.Second
	// WriteTimeout covers the time from the end of the request header read
	// to the end of the response write (a.k.a. the lifetime of the ServeHTTP)
	// When the connection is HTTPS, it also covers the packets written as
	// part of the TLS handshake and including the header read and the first byte wait.
	WriteTimeout = 1260 * time.Second
	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alive is enabled
	IdleTimeout = 600 * time.Second
	// HTTPClientTimeout specifies a time limit for requests made by a net/http Client.
	// The timeout includes connection time, any redirects, and reading the response body.
)

const (
	GoServerAppEndPoint = "http://localhost:8080"
)
