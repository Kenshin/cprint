package cprint

import (
	"errors"
	"testing"
)

func TestCP(t *testing.T) {

	// test one
	P(DEFAULT, "Test default color P() model", "\n")
	P(WARING, "Remote latest version %v %v latest version %v.\n", "0.10.28", "=", "0.1.0.26")
	P(DEFAULT, "\n")

	// test tow
	P(DEFAULT, "Test custom color P() model", "\n")
	cp := CP{Red, false, None, false, "="}
	P(NOTICE, "Remote latest version %v %v latest version %v, don't need to upgrade.\n", "0.10.28", cp, "0.1.0.26")
	P(DEFAULT, "\n")

	// test three
	P(DEFAULT, "Test Error() model", "\n")
	err := errors.New("Variable is not defined.")
	Error(ERROR, "'gnvm updte latest' an error has occurred. \nError: ", err)
	P(DEFAULT, "\n")
}
