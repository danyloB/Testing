# github.com/danyloB/Testing


	package suite // import "https://github.com/danyloB/Testing"
	
	Package suite implements an xUnit-style test runner, aiming for an optimum
	balance between simplicity and utility. It is based on the following
	packages:
	
	    - [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
	    - [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)
	
	For those using GoLand by JetBrains, you may find the following "live
	template" helpful:
	
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
	
	FUNCTIONS
	
	func Run(fixture interface{}, options ...Option)
	    Run accepts a fixture with Test* methods and optional setup/teardown methods
	    and executes the suite. Fixtures must be struct types which embed a
	    *testing.T. Assuming a fixture struct with test methods 'Test1' and 'Test2'
	    execution would proceed as follows:
	
	        1. fixture.SetupSuite()
	        2. fixture.Setup()
	        3. fixture.Test1()
	        4. fixture.Teardown()
	        5. fixture.Setup()
	        6. fixture.Test2()
	        7. fixture.Teardown()
	        8. fixture.TeardownSuite()
	
	    The methods provided by Options may be supplied to this function to tweak
	    the execution.
	
	
	TYPES
	
	type Opt struct{}
	
	var Options Opt
	    Options provides the sole entrypoint to the option functions provided by
	    this package.
	
	func (Opt) FreshFixture() Option
	    FreshFixture signals to Run that the new instances of the provided fixture
	    are to be instantiated for each and every test case. The Setup and Teardown
	    methods are also executed on the specifically instantiated fixtures. NOTE:
	    the SetupSuite and TeardownSuite methods are always run on the provided
	    fixture instance, regardless of this options having been provided.
	
	func (Opt) IntegrationTests() Option
	    IntegrationTests is a composite option that signals to Run that the test
	    suite should be treated as an integration test suite, avoiding parallelism
	    and utilizing shared fixtures to allow reuse of potentially expensive
	    resources.
	
	func (Opt) ParallelFixture() Option
	    ParallelFixture signals to Run that the provided fixture instance can be
	    executed in parallel with other go test functions. This option assumes that
	    `go test` was invoked with the -parallel flag.
	
	func (Opt) ParallelTests() Option
	    ParallelTests signals to Run that the test methods on the provided fixture
	    instance can be executed in parallel with each other. This option assumes
	    that `go test` was invoked with the -parallel flag.
	
	func (Opt) SharedFixture() Option
	    SharedFixture signals to Run that the provided fixture instance is to be
	    used to run all test methods. This mode is not compatible with
	    ParallelFixture or ParallelTests and disables them.
	
	func (Opt) UnitTests() Option
	    UnitTests is a composite option that signals to Run that the test suite can
	    be treated as a unit-test suite by employing parallelism and fresh fixtures
	    to maximize the chances of exposing unwanted coupling between tests.
	
	type Option func(*config)
	    Option is a function that modifies a config. See Options for provided
	    behaviors.
	
	type T struct{ *testing.T }
	    T embeds *testing.T and provides convenient hooks for making assertions and
	    other operations.
	
	func New(t *testing.T) *T
	    New prepares a *T for use with the fixture passed to Run.
	
	func (this *T) FatalSo(actual interface{}, assertion assertion, expected ...interface{}) bool
	    FatalSo is like So but in the event of an assertion failure it calls
	    *testing.T.Fatal.
	
	func (this *T) So(actual interface{}, assertion assertion, expected ...interface{}) bool
	    So invokes the provided assertion with the provided args. In the event of an
	    assertion failure it calls *testing.T.Error.
	
	func (this *T) Write(p []byte) (n int, err error)
	    Write implements io.Writer allowing for the suite to serve as a convenient
	    log target, among other use cases.
	
---

	package should // import "https://github.com/danyloB/Testing"
	
	
	VARIABLES
	
	var (
		ErrExpectedCountInvalid = errors.New("expected count invalid")
		ErrTypeMismatch         = errors.New("type mismatch")
		ErrKindMismatch         = errors.New("kind mismatch")
		ErrAssertionFailure     = errors.New("assertion failure")
	)
	var NOT negated
	    NOT (a singleton) constrains all negated assertions to their own namespace.
	
	
	FUNCTIONS
	
	func BeEmpty(actual interface{}, expected ...interface{}) error
	    BeEmpty uses reflection to verify that len(actual) == 0.
	
	func BeFalse(actual interface{}, expected ...interface{}) error
	    BeFalse verifies that actual is the boolean false value.
	
	func BeIn(actual interface{}, expected ...interface{}) error
	    BeIn determines whether actual is a member of expected[0]. It defers to
	    Contain.
	
	func BeNil(actual interface{}, expected ...interface{}) error
	    BeNil verifies that actual is the nil value.
	
	func BeTrue(actual interface{}, expected ...interface{}) error
	    BeTrue verifies that actual is the boolean true value.
	
	func Contain(actual interface{}, expected ...interface{}) error
	    Contain determines whether actual contains expected[0]. The actual value may
	    be a map, array, slice, or string:
	
	        - In the case of maps the expected value is assumed to be a map key.
	        - In the case of slices and arrays the expected value is assumed to be a member.
	        - In the case of strings the expected value may be a rune or substring.
	
	func EndWith(actual interface{}, expected ...interface{}) error
	    EndWith verifies that actual ends with expected[0]. The actual value may be
	    an array, slice, or string.
	
	func Equal(actual interface{}, EXPECTED ...interface{}) error
	    Equal verifies that the actual value is equal to the expected value. It uses
	    reflect.DeepEqual in most cases, but also compares numerics regardless of
	    specific type and compares time.Time values using the time.Equal method.
	
	func HaveLength(actual interface{}, expected ...interface{}) error
	    HaveLength uses reflection to verify that len(actual) == 0.
	
	func Panic(actual interface{}, expected ...interface{}) (err error)
	    Panic invokes the func() provided as actual and recovers from any panic. It
	    returns an error if actual() does not result in a panic.
	
	func StartWith(actual interface{}, expected ...interface{}) error
	    StartWith verified that actual starts with expected[0]. The actual value may
	    be an array, slice, or string.
	
	func WrapError(actual interface{}, expected ...interface{}) error
	    WrapError uses errors.Is to verify that actual is an error value that wraps
	    expected[0] (also an error value).
	

