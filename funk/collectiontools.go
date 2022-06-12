package funk

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

func AppendIfNotEQ[E any](sl []E, elem E, cmpfn func(a, b E) bool) []E {
	found := false
	for _, e := range sl {
		if cmpfn(e, elem) {
			found = true
		}
	}
	if found {
		return sl
	}
	return append(sl, elem)
}
