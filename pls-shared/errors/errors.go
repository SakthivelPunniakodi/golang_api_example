// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package errors

import (
	"fmt"
)

type Code uint

const (
	Unknown Code = iota
	NotFound
	InvalidArgument
)

type Error struct {
	orig error
	code Code
	msg  string
}

func NewErrorf(code Code, format string, args ...any) error {
	return WrapErrorf(nil, code, format, args...)
}

func WrapErrorf(err error, code Code, format string, args ...any) error {
	return &Error{
		orig: err,
		code: code,
		msg:  fmt.Sprintf(format, args...),
	}
}

func (e Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

func (e Error) Code() Code {
	return e.code
}
