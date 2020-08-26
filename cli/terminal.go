package cli

import (
	"fmt"
	"io"
	"os"
)

// Terminal implements terminal stdout
type Terminal struct{}

// Out stdout
func (*Terminal) Out() io.Writer {
	return os.Stdout
}

// In stdin
func (*Terminal) In() io.Reader {
	return os.Stdin
}

// Printf use fmt printf
func (*Terminal) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
