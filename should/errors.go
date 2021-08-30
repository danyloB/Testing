package should

import (
	"errors"
	"fmt"
)

var (
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

func failure(format string, args ...interface{}) error {
	return wrap(ErrAssertionFailure, format, args...)
}
func wrap(inner error, format string, args ...interface{}) error {
	return fmt.Errorf("%w: "+fmt.Sprintf(format, args...), inner)
}
