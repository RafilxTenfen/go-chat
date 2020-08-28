package cli

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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

// ReadPassword reads a password from terminal
func (t *Terminal) ReadPassword() string {
	// initial state of the terminal
	inittermstate, err := terminal.GetState(syscall.Stdin)
	if err != nil {
		fmt.Println(err)
	}

	// restore it in the event of an interrupt
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, inittermstate)
		os.Exit(1)
	}()

	// getpass
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		fmt.Println(err)
	}
	// looking for ^C
	signal.Stop(c)

	return string(p)
}
