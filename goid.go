// Package goid provides runtime info of current goroutine.
package goid

// GoID returns the current goroutine id.
// It exactly matches goroutine id of the stack trace.
// Note that the user should not use this id to implement feature such as goroutine local storage.
// It is intended to support debug logger or testing.
func GoID() int64
