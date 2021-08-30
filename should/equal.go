package should

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases, but also compares numerics
// regardless of specific type and compares time.Time values using the
// time.Equal method.
func Equal(actual interface{}, EXPECTED ...interface{}) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}

	expected := EXPECTED[0]

	for _, spec := range specs {
		if !spec.isSatisfiedBy(actual, expected) {
			continue
		}
		if spec.areEqual(actual, expected) {
			return nil
		}
		break
	}
	return failure(report(actual, expected))
}

// Equal negated!
func (negated) Equal(actual interface{}, expected ...interface{}) error {
	err := Equal(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		expected[0],
		actual,
	)
}

var specs = []specification{
	numericEquality{},
	timeEquality{},
	deepEquality{},
}

func report(a, b interface{}) string {
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := fmt.Sprintf(format(a), a)
	bFormat := fmt.Sprintf(format(b), b)
	typeDiff := diff(bType, aType)
	valueDiff := diff(bFormat, aFormat)

	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "Expected: %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "Actual  : %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "          %s %s\n", typeDiff, valueDiff)
	_, _ = fmt.Fprintf(builder, "Stack (filtered):\n%s\n", stack())

	return builder.String()
}
func format(v interface{}) string {
	if isNumeric(v) || isTime(v) {
		return "%v"
	} else {
		return "%#v"
	}
}
func diff(a, b string) string {
	result := new(strings.Builder)

	for x := 0; ; x++ {
		if x >= len(a) && x >= len(b) {
			break
		}
		if x >= len(a) || x >= len(b) || a[x] != b[x] {
			result.WriteString("^")
		} else {
			result.WriteString(" ")
		}
	}
	return result.String()
}
func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		if strings.Contains(lines[x+1], "_test.go:") {
			filtered = append(filtered, lines[x], lines[x+1])
		}
	}
	return "> " + strings.Join(filtered, "\n> ")
}

type formatter func(interface{}) string

type specification interface {
	isSatisfiedBy(a, b interface{}) bool
	areEqual(a, b interface{}) bool
}

// deepEquality compares any two values using reflect.DeepEqual.
// https://golang.org/pkg/reflect/#DeepEqual
type deepEquality struct{}

func (this deepEquality) isSatisfiedBy(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (this deepEquality) areEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// numericEquality compares numeric values using the built-in equality
// operator (`==`). Values of differing numeric reflect.Kind are each
// converted to the type of the other and are compared with `==` in both
// directions. https://golang.org/pkg/reflect/#Kind
type numericEquality struct{}

func (this numericEquality) isSatisfiedBy(a, b interface{}) bool {
	return isNumeric(a) && isNumeric(b)
}
func (this numericEquality) areEqual(a, b interface{}) bool {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	aAsB := aValue.Convert(bValue.Type()).Interface()
	bAsA := bValue.Convert(aValue.Type()).Interface()
	return a == bAsA && b == aAsB
}
func isNumeric(v interface{}) bool {
	_, found := numericKinds[reflect.TypeOf(v).Kind()]
	return found
}

// timeEquality compares values both of type time.Time using their Equal method.
// https://golang.org/pkg/time/#Time.Equal
type timeEquality struct{}

func (this timeEquality) isSatisfiedBy(a, b interface{}) bool {
	return isTime(a) && isTime(b)
}
func (this timeEquality) areEqual(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}
func isTime(v interface{}) bool {
	_, ok := v.(time.Time)
	return ok
}

var numericKinds = map[reflect.Kind]struct{}{
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
}
