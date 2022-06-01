package assertor

import "testing"

func AssertEquals(t *testing.T, expected any, got any) {

	if got != expected {
		t.Fatalf("expected %v; got %v", expected, got)
	}
}
