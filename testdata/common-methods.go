// Test that we don't nag for comments on common methods.
// OK

// Package pkg ...
package pkg

// T is ...
type T int

func (T) Error() string                     { return "" }
func (T) String() string                    { return "" }
func (T) Read(p []byte) (n int, err error)  { return 0, nil }
func (T) Write(p []byte) (n int, err error) { return 0, nil }
