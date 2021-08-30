package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeFalse(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeFalse, "EXTRA")
	assert.TypeMismatch(1, should.BeFalse)

	assert.Fail(true, should.BeFalse)
	assert.Pass(false, should.BeFalse)
}
