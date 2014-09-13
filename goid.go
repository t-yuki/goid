package goid

// ID returns the current goroutine id.
// It exactly matches goroutine id of the stack trace.
// Note that the user should not use this id to implement feature such as goroutine local storage.
// It is intended to support debug logger or testing.
func ID() int32
