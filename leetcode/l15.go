package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	l := len(nums)
	if l < 3 {
		return result
	}
	if nums[0] > 0 {
		return result
	}
	if nums[l-1] < 0 {
		return result
	}
	sort.Ints(nums)
	fmt.Println("-----------------")

	for k := 0; k < l; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		target := 0 - nums[k]
		i := k + 1
		j := l + 1
		for i < j {
			if nums[i]+nums[j] == target {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i = i + 1
				}
				for i < j && nums[j] == nums[j-1] {
					j = j - 1
				}
				i = i + 1
				j = j - 1
			} else if nums[i]+nums[j] < target {
				i = i + 1
			} else {
				j = j - 1
			}

		}
	}
	return result
}
