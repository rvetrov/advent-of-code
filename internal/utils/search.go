package utils

func LowerBound(left, right int, isOnRight func(int) bool) int {
	for left < right {
		mid := (left + right) / 2
		if isOnRight(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
