package gofunk

// MapSlice does a map() on a Slice; producing a Slice
// #(a b c) collect: fn
func MapSlice[E any, R any](sl []E, fn func(E) R) []R {
	coll := make([]R, len(sl))
	for i, e := range sl {
		coll[i] = fn(e)
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
