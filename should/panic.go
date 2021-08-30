package should

import "errors"

// Panic invokes the func() provided as actual and recovers from any
// panic. It returns an error if actual() does not result in a panic.
func Panic(actual interface{}, expected ...interface{}) (err error) {
	err = NOT.Panic(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("" +
		"provided func did not panic as expected " +
		"(...or it panicked with a <nil> value...)",
	)
}

// Panic (negated!) expects the func() provided as actual to run without panicking.
func (negated) Panic(actual interface{}, expected ...interface{}) (err error) {
	err = validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, func() {})
	if err != nil {
		return err
	}

	defer func() {
		r := recover()
		if r != nil {
			err = failure(""+
				"provided func should not have"+
				"panicked but it did with: %s", r,
			)
		}
	}()

	actual.(func())()
	return nil
}
