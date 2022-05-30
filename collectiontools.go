package gofunk

// #(a b c) collect: fn
func MapSlice[E any, R any](sl []E, fn func(E) R) []R {
	coll := make([]R, len(sl))
	for i, e := range sl {
		coll[i] = fn(e)
	}
	return coll
}

// (1 to: 100) collect: fn
func MapRange[R any](start int, stop int, fn func(int) R) []R {
	coll := make([]R, stop-start)
	for i := start; i < stop; i++ {
		coll[i] = fn(i)
	}
	return coll
}
