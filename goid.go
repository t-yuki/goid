// Package goid provides a function that gets goroutine id of current goroutine.
// Run `go generate` before use.
package goid

//go:generate sh gen.sh $GOROOT

// GoID returns the current goroutine id.
// It exactly matches goroutine id of the stack trace.
// Note that users SHOULD NOT use this id to implement functional features such as goroutine local storage.
// It is intended to support debug logger or testing.
func GoID() int64
