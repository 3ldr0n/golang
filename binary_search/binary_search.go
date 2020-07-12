package binary_search

func binary_search(x int, list []int) bool {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == x {
			return true
		} else if list[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
