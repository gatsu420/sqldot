package tests

import (
	"reflect"
	"testing"
)

func AssertEqualObject(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("had %#v \n of type %v \n, expected %#v \n of type %v \n",
			actual,
			reflect.TypeOf(actual),
			expected,
			reflect.TypeOf(expected),
		)
		return
	}
}

func AssertEqualError(t *testing.T, actual error, expected error) {
	if actual == nil || expected == nil {
		if actual != expected {
			t.Errorf("had %v, expected %v", actual, expected)
		}
		return
	}
	if actual.Error() != expected.Error() {
		t.Errorf("had %v, expected %v", actual.Error(), expected.Error())
		return
	}
}
