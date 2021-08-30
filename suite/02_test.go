package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: suite.New(t)}
	suite.Run(fixture, suite.Options.UnitTests())
	fixture.So(fixture.counter, should.Equal, 0)
}

type Suite02 struct {
	*suite.T
	counter int
}

func (this *Suite02) TestSomething() {
	_, _ = this.Write([]byte("*** this should appear in the test log!"))
	this.counter++
}
