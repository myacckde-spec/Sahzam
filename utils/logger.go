package utils

import "fmt"

// Logger writes simple educational progress messages.
type Logger interface {
	Log(msg string)
}

// StdLogger prints messages to stdout.
type StdLogger struct{}

// Log prints a message.
func (l *StdLogger) Log(msg string) {
	fmt.Println(msg)
}
