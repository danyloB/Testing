package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite06{T: suite.New(t)}

	suite.Run(fixture, suite.Options.SharedFixture())

	fixture.So(fixture.events, should.BeNil)
}

type Suite06 struct {
	*suite.T
	events []string
}

func (this *Suite06) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite06) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) SkipTest()           { this.record("SkipTest") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }
