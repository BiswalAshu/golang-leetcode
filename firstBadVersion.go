package main

// givien API bool isBadVersion(int)
// minimize calls

func firstBadVersion(n int) int {
	first := 1
	last := n
	var bad int
	for first <= last {
		mid := (first + last) / 2
		if isBadVersion(mid) {
			bad = mid
			last = mid - 1
		} else {
			bad = mid
			first = mid + 1
		}
	}

	return bad
}
