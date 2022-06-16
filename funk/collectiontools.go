package funk

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// MapSlice does a map() on a Slice; producing a Slice
// #(a b c) collect: fn
func MapSlice[E any, R any](sl []E, fn func(E) R) []R {
	coll := make([]R, len(sl))
	for i, e := range sl {
		coll[i] = fn(e)
	}
	return coll
}

// MapKVMap does a map() on a Map; producing a Map
func MapKVMap[K comparable, E any, R any](mp map[K]E, fn func(K, E) R) map[K]R {
	coll := make(map[K]R, len(mp))
	for k, e := range mp {
		coll[k] = fn(k, e)
	}
	return coll
}

// MapMap does a map() on a Map; producing a Map
func MapMap[K comparable, E any, R any](mp map[K]E, fn func(E) R) map[K]R {
	coll := make(map[K]R, len(mp))
	for k, e := range mp {
		coll[k] = fn(e)
	}
	return coll
}

func MapToSlice[K constraints.Ordered, E any, R any](mp map[K]E, fn func(E) R) []R {
	keys := make([]K, 0)
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	coll := make([]R, 0, len(keys))
	for _, k := range keys {
		coll = append(coll, fn(mp[k]))
	}
	return coll
}

func MapKVToSlice[K constraints.Ordered, E any, R any](mp map[K]E, fn func(K, E) R) []R {
	keys := make([]K, 0)
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	coll := make([]R, 0, len(keys))
	for _, k := range keys {
		coll = append(coll, fn(k, mp[k]))
	}
	return coll
}

// MapRange does a map on all integers between start and stop; producing a Slice
// (1 to: 100) collect: fn
func MapRange[R any](start int, stop int, fn func(int) R) []R {
	coll := make([]R, stop-start)
	for i := start; i < stop; i++ {
		coll[i-start] = fn(i)
	}
	return coll
}

// FilterSlice does a filter() on a Slice; producing a Slice
// #(a b c) select: fn
func FilterSlice[E any](sl []E, fn func(E) bool) []E {
	coll := make([]E, 0)
	for _, e := range sl {
		if fn(e) {
			coll = append(coll, e)
		}
	}
	return coll
}

func SoleElement[E any](sl []E) E {
	if len(sl) == 1 {
		return sl[0]
	}
	if len(sl) == 0 {
		panic("No element found")
	}
	panic("Multiple elements found")
}

func AppendIfAbsent[E any](sl []E, elem E) []E {
	found := false
	for _, e := range sl {
		if interface{}(e) == interface{}(elem) {
			found = true
		}
	}
	if found {
		return sl
	}
	return append(sl, elem)
}

func AppendIfNotEQ[E any](sl []E, elem E, cmp func(a, b E) bool) []E {
	found := false
	for _, e := range sl {
		if cmp(e, elem) {
			found = true
		}
	}
	if found {
		return sl
	}
	return append(sl, elem)
}

func CopySlice[E any](sl []E) []E {
	var cp []E
	copy(cp, sl)
	return cp
}

func ConcatSlices[E any](slices ...[]E) []E {
	var cp []E
	for _, slice := range slices {
		cp = append(cp, slice...)
	}
	return cp
}

func EqualsMap[K comparable, V any](mapA map[K]V, mapB map[K]V, cmp func(a, b V) bool) bool {

	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapA {
		ov, found := mapB[k]
		if !found {
			return false
		}
		if !cmp(v, ov) {
			return false
		}
	}
	return true
}

func EqualsSlice[V any](sliceA []V, sliceB []V, cmp func(a, b V) bool) bool {

	if len(sliceA) != len(sliceB) {
		return false
	}

	for idx, v := range sliceA {
		ov := sliceB[idx]
		if !cmp(v, ov) {
			return false
		}
	}
	return true
}

func SliceDetect[V any](sl []V, test func(e V) bool) (*V, bool) {
	for _, v := range sl {
		if test(v) {
			return &v, true
		}
	}
	return nil, false
}

func SliceAnySatisfy[V any](sl []V, test func(e V) bool) bool {
	for _, v := range sl {
		if test(v) {
			return true
		}
	}
	return false
}

func MapAnySatisfy[K comparable, V any](mp map[K]V, test func(e V) bool) bool {
	for _, v := range mp {
		if test(v) {
			return true
		}
	}
	return false
}

func FoldSlice[V any](sl []V, fold func(a, b V) V) V {
	if len(sl) == 0 {
		var noop V
		return noop
	}
	firstFold := true
	var val V
	for _, v := range sl {
		if firstFold {
			val = v
			firstFold = false
		} else {
			val = fold(val, v)
		}
	}
	return val
}
