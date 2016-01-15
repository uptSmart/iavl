package types

import (
	"errors"
)

type RetCode int

// Reserved return codes
const (
	RetCodeOK            RetCode = 0
	RetCodeInternalError RetCode = 1
)

func (r RetCode) Error() error {
	switch r {
	case RetCodeOK:
		return nil
	default:
		return errors.New(r.String())
	}
}

//go:generate stringer -type=RetCode

// NOTE: The previous comment generates r.String().
// To run it, `go get golang.org/x/tools/cmd/stringer`
// and `go generate` in tmsp/types
