package assert_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/mdwhatcott/testing/assert"
)

func TestSo(t *testing.T) {
	assertNil(t, assert.So(1, shouldPass))
	assertErr(t, assert.So(1, shouldFail))
}

func TestLog_Pass_Nop(t *testing.T) {
	fakeT := new(FakeT)

	assert.Log(fakeT).So(1, shouldPass)
	assert.Log(fakeT).So(1, shouldPass)
	assert.Log(fakeT).So(1, shouldPass)

	assertEqual(t, fakeT, new(FakeT))
}
func TestLog_Fail_Logs(t *testing.T) {
	fakeT := new(FakeT)

	assert.Log(fakeT).So(1, shouldFail)
	assert.Log(fakeT).So(1, shouldFail)
	assert.Log(fakeT).So(1, shouldFail)

	assertEqual(t, fakeT, &FakeT{
		helps:  3,
		logs:   []string{"failure", "failure", "failure"},
		errors: nil,
		fatals: nil,
	})
}

func TestError_Pass_Nop(t *testing.T) {
	fakeT := new(FakeT)

	assert.Error(fakeT).So(1, shouldPass)
	assert.Error(fakeT).So(1, shouldPass)
	assert.Error(fakeT).So(1, shouldPass)

	assertEqual(t, fakeT, new(FakeT))
}
func TestError_Fail_Errors(t *testing.T) {
	fakeT := new(FakeT)

	assert.Error(fakeT).So(1, shouldFail)
	assert.Error(fakeT).So(1, shouldFail)
	assert.Error(fakeT).So(1, shouldFail)

	assertEqual(t, fakeT, &FakeT{
		helps:  3,
		logs:   nil,
		errors: []string{"failure", "failure", "failure"},
		fatals: nil,
	})
}

func TestFatal_Pass_Nop(t *testing.T) {
	fakeT := new(FakeT)

	assert.Fatal(fakeT).So(1, shouldPass)
	assert.Fatal(fakeT).So(1, shouldPass)
	assert.Fatal(fakeT).So(1, shouldPass)

	assertEqual(t, fakeT, new(FakeT))
}
func TestFatal_Fail_Fatals(t *testing.T) {
	fakeT := new(FakeT)

	assert.Fatal(fakeT).So(1, shouldFail)
	assert.Fatal(fakeT).So(1, shouldFail)
	assert.Fatal(fakeT).So(1, shouldFail)

	assertEqual(t, fakeT, &FakeT{
		helps:  3,
		logs:   nil,
		errors: nil,
		fatals: []string{"failure", "failure", "failure"},
	})
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		return
	}
	t.Errorf("\n"+
		"expected: %#v\n"+
		"actual:   %#v",
		expected,
		actual,
	)
}
func assertErr(t *testing.T, err error) {
	if err != nil {
		return
	}
	t.Helper()
	t.Error("Expected non-<nil> value, got:", err)
}
func assertNil(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Helper()
	t.Error("Expected <nil> value, got:", err)
}

func shouldPass(actual interface{}, expected ...interface{}) error {
	_ = actual
	_ = expected
	return nil
}
func shouldFail(actual interface{}, expected ...interface{}) error {
	_ = actual
	_ = expected
	return errors.New("failure")
}

type FakeT struct {
	helps  int
	logs   []string
	errors []string
	fatals []string
}

func (this *FakeT) Helper() { this.helps++ }

func (this *FakeT) Log(args ...interface{}) {
	this.logs = append(this.logs, fmt.Sprint(args...))
}
func (this *FakeT) Error(args ...interface{}) {
	this.errors = append(this.errors, fmt.Sprint(args...))
}
func (this *FakeT) Fatal(args ...interface{}) {
	this.fatals = append(this.fatals, fmt.Sprint(args...))
}
