package should

import (
	"errors"
	"reflect"
)

// BeEmpty uses reflection to verify that len(actual) == 0.
func BeEmpty(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, kindsWithLength...)
	if err != nil {
		return err
	}

	length := reflect.ValueOf(actual).Len()
	if length == 0 {
		return nil
	}

	TYPE := reflect.TypeOf(actual).String()
	return failure("got len(%s) == %d, want empty %s", TYPE, length, TYPE)
}

// BeEmpty (negated!)
func (negated) BeEmpty(actual interface{}, expected ...interface{}) error {
	err := BeEmpty(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}
	if err != nil {
		return err
	}
	TYPE := reflect.TypeOf(actual).String()
	return failure("got empty %s, want non-empty %s", TYPE, TYPE)
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}
