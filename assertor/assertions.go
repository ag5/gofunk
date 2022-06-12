package assertor

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expected any, got any) {

	if got != expected {
		t.Fatalf("expected %v; got %v", expected, got)
	}
}

func AssertDeepEquals(t *testing.T, expected any, got any) {
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v; got %v", expected, got)
	}
}
