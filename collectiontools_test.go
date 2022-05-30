package gofunk

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapSlice(t *testing.T) {
	src := []string{"foo", "bar", "zonk"}

	got := MapSlice(src, func(st string) int { return len(st) })

	expected := []int{3, 3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("%v expected; got %v", expected, got)
	}
}

func TestMapRange(t *testing.T) {

	got := MapRange(1, 4, func(i int) string { return strconv.Itoa(i) })

	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("%v expected; got %v", expected, got)
	}
}

func TestFilterSlice(t *testing.T) {
	src := []string{"foo", "bar", "zonk"}

	got := FilterSlice(src, func(st string) bool { return len(st) == 4 })

	expected := []string{"zonk"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("%v expected; got %v", expected, got)
	}
}
