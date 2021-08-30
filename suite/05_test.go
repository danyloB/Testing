package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestFocus(t *testing.T) {
	fixture := &Suite05{
		T:      suite.New(t),
		events: make(map[string]struct{}),
	}

	suite.Run(fixture, suite.Options.SharedFixture())

	fixture.So(t.Failed(), should.BeFalse)
	if testing.Short() {
		fixture.So(fixture.events, should.Equal, map[string]struct{}{"1": {}})
	} else {
		fixture.So(fixture.events, should.Equal, map[string]struct{}{
			"1": {},
			"2": {},
		})
	}
}

type Suite05 struct {
	*suite.T
	events map[string]struct{}
}

func (this *Suite05) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite05) FocusLongTest2() {
	this.events["2"] = struct{}{}
}
func (this *Suite05) TestThatFails() {
	this.So(1, should.Equal, 2)
}
