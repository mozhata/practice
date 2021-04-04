package basic

func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/*
搜索插入位置
https://leetcode-cn.com/problems/search-insert-position/
*/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	idx := len(nums)
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
			idx = mid
		} else {
			l = mid + 1
		}
	}
	return idx
}
