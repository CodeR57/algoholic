package algoholic

func binarySearchImpl(ns []int, n, offset int) int {
	if len(ns) == 0 {
		return -1
	}

	half := len(ns) / 2
	val := ns[half]

	switch {
	case n == val:
		return half + offset
	case n < val:
		return binarySearchImpl(ns[:half], n, offset)
	default:
		// n > val
		return binarySearchImpl(ns[half+1:], n, offset+half+1)
	}
}

func BinarySearchRecursive(ns []int, n int) int {
	return binarySearchImpl(ns, n, 0)
}

func BinarySearchIterative(ns []int, n int) int {
	if len(ns) == 0 {
		return -1
	}

	from, to := 0, len(ns)-1

	for {
		if from > to {
			return -1
		}

		// Avoid overflow.
		mid := to - (to-from)/2
		val := ns[mid]

		switch {
		case val == n:
			return mid
		case n < val:
			to = mid - 1
		default:
			from = mid + 1
		}
	}
}
