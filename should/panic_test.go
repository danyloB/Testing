package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.Panic)

	assert.Fail(func() {}, should.Panic)
	assert.Pass(func() { panic("yay") }, should.Panic)
}

func TestShouldNotPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.NOT.Panic)

	assert.Fail(func() { panic("boo") }, should.NOT.Panic)
	assert.Pass(func() {}, should.NOT.Panic)
}
