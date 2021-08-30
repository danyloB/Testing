/*
Package suite implements an xUnit-style test
runner, aiming for an optimum balance between
simplicity and utility. It is based on the
following packages:

	- [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
	- [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)

For those using GoLand by JetBrains, you may
find the following "live template" helpful:

	func Test$NAME$Suite(t *testing.T) {
		suite.Run(&$NAME$Suite{T: suite.New(t)}, suite.Options.UnitTests())
	}

	type $NAME$Suite struct {
		*suite.T
	}

	func (this *$NAME$Suite) Setup() {
	}

	func (this *$NAME$Suite) Test$END$() {
	}

Happy testing!
*/
package suite

import (
	"reflect"
	"strings"
	"testing"
)

/*
Run accepts a fixture with Test* methods and
optional setup/teardown methods and executes
the suite. Fixtures must be struct types which
embed a *testing.T. Assuming a fixture struct
with test methods 'Test1' and 'Test2' execution
would proceed as follows:

	1. fixture.SetupSuite()
	2. fixture.Setup()
	3. fixture.Test1()
	4. fixture.Teardown()
	5. fixture.Setup()
	6. fixture.Test2()
	7. fixture.Teardown()
	8. fixture.TeardownSuite()

The methods provided by Options may be supplied
to this function to tweak the execution.
*/
func Run(fixture interface{}, options ...Option) {
	config := new(config)
	for _, option := range options {
		option(config)
	}

	fixtureValue := reflect.ValueOf(fixture)
	fixtureType := reflect.TypeOf(fixture)
	t := fixtureValue.Elem().FieldByName("T").Interface().(*T)

	var (
		testNames        []string
		skippedTestNames []string
		focusedTestNames []string
	)
	for x := 0; x < fixtureType.NumMethod(); x++ {
		name := fixtureType.Method(x).Name
		method := fixtureValue.MethodByName(name)
		_, isNiladic := method.Interface().(func())
		if !isNiladic {
			continue
		}

		if strings.HasPrefix(name, "Test") {
			testNames = append(testNames, name)
		} else if strings.HasPrefix(name, "LongTest") {
			testNames = append(testNames, name)

		} else if strings.HasPrefix(name, "SkipLongTest") {
			skippedTestNames = append(skippedTestNames, name)
		} else if strings.HasPrefix(name, "SkipTest") {
			skippedTestNames = append(skippedTestNames, name)

		} else if strings.HasPrefix(name, "FocusLongTest") {
			focusedTestNames = append(focusedTestNames, name)
		} else if strings.HasPrefix(name, "FocusTest") {
			focusedTestNames = append(focusedTestNames, name)
		}
	}

	if len(focusedTestNames) > 0 {
		testNames = focusedTestNames
	}

	if len(testNames) == 0 {
		t.Skip("NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)")
		return
	}

	if config.parallelFixture {
		t.Parallel()
	}

	setup, hasSetup := fixture.(setupSuite)
	if hasSetup {
		setup.SetupSuite()
	}

	teardown, hasTeardown := fixture.(teardownSuite)
	if hasTeardown {
		defer teardown.TeardownSuite()
	}

	for _, name := range skippedTestNames {
		testCase{t: t, manualSkip: true, name: name}.Run()
	}

	for _, name := range testNames {
		testCase{t, name, config, false, fixtureType, fixtureValue}.Run()
	}
}

type testCase struct {
	t            *T
	name         string
	config       *config
	manualSkip   bool
	fixtureType  reflect.Type
	fixtureValue reflect.Value
}

func (this testCase) Run() {
	_ = this.t.Run(this.name, this.decideRun())
}
func (this testCase) decideRun() func(*testing.T) {
	if this.manualSkip {
		return this.skipFunc("Skipping: " + this.name)
	}

	if isLongRunning(this.name) && testing.Short() {
		return this.skipFunc("Skipping long-running test in -test.short mode: " + this.name)
	}

	return this.runTest
}
func (this testCase) skipFunc(message string) func(*testing.T) {
	return func(t *testing.T) { t.Skip(message) }
}
func (this testCase) runTest(t *testing.T) {
	if this.config.parallelTests {
		t.Parallel()
	}

	fixtureValue := this.fixtureValue
	if this.config.freshFixture {
		fixtureValue = reflect.New(this.fixtureType.Elem())
	}
	fixtureValue.Elem().FieldByName("T").Set(reflect.ValueOf(&T{T: t}))

	setup, hasSetup := fixtureValue.Interface().(setupTest)
	if hasSetup {
		setup.Setup()
	}

	teardown, hasTeardown := fixtureValue.Interface().(teardownTest)
	if hasTeardown {
		defer teardown.Teardown()
	}

	fixtureValue.MethodByName(this.name).Call(nil)
}

func isLongRunning(name string) bool {
	return strings.HasPrefix(name, "Long") ||
		strings.HasPrefix(name, "FocusLong")
}

type (
	setupSuite    interface{ SetupSuite() }
	setupTest     interface{ Setup() }
	teardownTest  interface{ Teardown() }
	teardownSuite interface{ TeardownSuite() }
)
