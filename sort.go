package slices

// insertionSort sorts v[begin:end) using insertion sort.
func insertionSort[T any](v []T, less func(a, b T) bool) {
	for cur := 1; cur < len(v); cur++ {
		for j := cur; j > 0 && less(v[j], v[j-1]); j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

// siftDown implements the heap property on v[lo:hi].
func siftDown[T any](v []T, less func(a, b T) bool, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(v[child], v[child+1]) {
			child++
		}
		if !less(v[root], v[child]) {
			return
		}
		v[root], v[child] = v[child], v[root]
		root = child
	}
}

func heapSort[T any](v []T, less func(a, b T) bool) {
	lo := 0
	hi := len(v)

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(v, less, i, hi)
	}

	// Pop elements into end of v.
	for i := hi - 1; i >= 0; i-- {
		v[0], v[i] = v[i], v[0]
		siftDown(v, less, lo, i)
	}
}
