package funk

import (
	"github.com/stretchr/testify/require"
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

func TestMapKVMap(t *testing.T) {
	src := map[string]string{"foo": "bar", "baz": "frob"}
	got := MapKVMap(src, func(k string, v string) int { return len(v) })

	expected := map[string]int{"foo": 3, "baz": 4}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("%v expected; got %v", expected, got)
	}
}

func TestMapMap(t *testing.T) {
	src := map[string]string{"foo": "bar", "baz": "frob"}
	got := MapMap(src, func(v string) int { return len(v) })

	expected := map[string]int{"foo": 3, "baz": 4}
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

func TestAppendIfAbsent(t *testing.T) {
	src := []string{"foo", "bar", "zonk"}
	src = AppendIfAbsent(src, "foo")

	var expected any = []string{"foo", "bar", "zonk"}
	var got any = src
	require.Equal(t, expected, got)

	src = AppendIfAbsent(src, "boo")

	var expected2 any = []string{"foo", "bar", "zonk", "boo"}
	var got2 any = src
	require.Equal(t, expected2, got2)

}

func TestAppendIfNotEQ(t *testing.T) {
	src := []string{"foo", "zonk"}
	src = AppendIfNotEQ(src, "bar", func(a, b string) bool { return len(a) == len(b) })

	var expected any = []string{"foo", "zonk"}
	var got any = src
	require.Equal(t, expected, got)

	src = AppendIfNotEQ(src, "frobnicate", func(a, b string) bool { return len(a) == len(b) })

	var expected2 any = []string{"foo", "zonk", "frobnicate"}
	var got2 any = src
	require.Equal(t, expected2, got2)
}
