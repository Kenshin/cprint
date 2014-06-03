package cprint

import (
	"errors"
	"testing"
)

func TestCP(t *testing.T) {

	// test ont
	P(WARING, "Remote latest version %v %v latest version %v.\n", "0.10.28", "=", "0.1.0.26")

	// test tow
	cp := CP{Red, false, None, false, "="}
	P(NOTICE, "Remote latest version %v %v latest version %v, don't need to upgrade.\n", "0.10.28", cp, "0.1.0.26")

	// test three
	err := errors.New("Variable is not defined.")
	Error(ERROR, "'gnvm updte latest' an error has occurred. \nError: ", err)
}
