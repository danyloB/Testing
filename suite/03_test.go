package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: suite.New(t)}
	suite.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite03 struct{ *suite.T }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, should.Equal, 2)
}
func (this *Suite03) SkipLongTestThatFails() {
	this.So(1, should.Equal, 2)
}
