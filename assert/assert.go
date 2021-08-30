package assert

type Assertion func(actual interface{}, expected ...interface{}) error

// So runs the provided Assertion and returns the error, as in:
// err := assert.So(1, should.Equal, 1)
func So(actual interface{}, assertion Assertion, expected ...interface{}) error {
	return assertion(actual, expected...)
}

type testingT interface {
	Helper()
	Log(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// Log receives a *testing.T, and prepares the caller for a So call.
// In the event of an assertion failure it will pass the err to *testing.T.Log.
// assert.Log(t).So(1, should.Equal, 2) // results in t.Log(err)
func Log(t testingT) TestingT { return TestingT{helper: t.Helper, report: t.Log} }

// Error receives a *testing.T, and prepares the caller for a So call.
// In the event of an assertion failure it will pass the err to *testing.T.Error.
// assert.Error(t).So(1, should.Equal, 2) // results in t.Error(err)
func Error(t testingT) TestingT { return TestingT{helper: t.Helper, report: t.Error} }

// Fatal receives a *testing.T, and prepares the caller for a So call.
// In the event of an assertion failure it will pass the err to *testing.T.Fatal.
// assert.Fatal(t).So(1, should.Equal, 2) // results in t.Fatal(err)
func Fatal(t testingT) TestingT { return TestingT{helper: t.Helper, report: t.Fatal} }

// TestingT is an intermediate type, not for direct instantiation.
type TestingT struct {
	helper func()
	report func(...interface{})
}

// So runs the provided Assertion and calls the configured reporting function, as in:
// - assert.Log(t).So(1, should.Equal, 1)
// - assert.Error(t).So(1, should.Equal, 1)
// - assert.Fatal(t).So(1, should.Equal, 1)
func (this TestingT) So(actual interface{}, assertion Assertion, expected ...interface{}) {
	err := assertion(actual, expected...)
	if err != nil {
		this.helper()
		this.report(err)
	}
}
