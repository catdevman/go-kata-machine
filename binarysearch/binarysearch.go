package binarysearch

import (
    "golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](nums []T, target T) int{
    return recursiveSearch(nums, 0, len(nums)-1, target)
    // return iterativeSearch(nums, target)
}

func recursiveSearch[T constraints.Ordered](nums []T, left int, right int, target T) int{
	var middle int
    if right >= left {
		middle = left + (right-left)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			return recursiveSearch(nums, left, middle-1, target)
		}
		return recursiveSearch(nums, middle+1, right, target)
	}
	return -1
}

func iterativeSearch[T constraints.Ordered](nums []T, target T) int{
    var left, middle, right int
	left, right = 0, len(nums)-1
	for right >= left {
		middle = left + (right-left)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			right = middle - 1
			continue
		}
		left = middle + 1
	}
	return -1
}
