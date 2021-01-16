package helpers

import (
	"reflect"
	"testing"
)

// AssertEquals fails the tests and show an error message if expected and got arguments
// are not equals.
func AssertEquals(t *testing.T, expected, got interface{}) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}

// AssertNoError fails the tests and show an error message if receives an error.
func AssertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("error not expected but %v", got)
	}
}

// AssertError fails the tests and show an error message if expected and got arguements
// are not equal errors.
func AssertError(t *testing.T, expected, got error) {
	t.Helper()

	if got == nil {
		t.Errorf("an error was expected but none given")
	}

	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}

// AssertDeepEquals has a similar behaviour to AssertEquals but performs the comparissons with the
// built in function reflection.DeepEquals (whose perfomance is worst) instead of the equality operator
func AssertDeepEquals(t *testing.T, expected, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
