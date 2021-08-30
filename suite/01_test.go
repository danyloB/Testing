package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestSuiteWithSetupsAndTeardowns(t *testing.T) {
	fixture := &Suite01{T: suite.New(t)}

	suite.Run(fixture, suite.Options.IntegrationTests())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite01 struct {
	*suite.T
	events []string
}

func (this *Suite01) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite01) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite01) Setup()              { this.record("Setup") }
func (this *Suite01) Teardown()           { this.record("Teardown") }
func (this *Suite01) Test()               { this.record("Test") }
func (this *Suite01) record(event string) { this.events = append(this.events, event) }
