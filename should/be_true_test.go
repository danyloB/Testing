package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeTrue(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeTrue, "EXTRA")

	assert.TypeMismatch(1, should.BeTrue)

	assert.Fail(false, should.BeTrue)
	assert.Pass(true, should.BeTrue)
}
