package binarysearch

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int

	getit := false
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			getit = true
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if getit {
		for mid > 0 && nums[mid-1] == target {
			mid -= 1
		}
		return mid

	}

	return -1
}
