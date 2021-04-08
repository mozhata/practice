package basic

/*
十大排序算法 https://www.cnblogs.com/onepixel/p/7674659.html
https://leetcode-cn.com/problems/sort-an-array/solution/pai-xu-shu-zu-by-leetcode-solution/
https://leetcode-cn.com/problems/sort-an-array/solution/golang-by-xilepeng-2/
*/
func BubbleSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func QuickSort(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(arr []int) {
	n := len(arr)
	left, right := 0, n-1
	// 滑动窗口, 直到左边大于P右边小于P, 做一次交换
	for right > left {
		if arr[right] > arr[0] {
			right--
			continue
		}
		if arr[left] <= arr[0] {
			left++
			continue
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	// 最终left right紧挨着 且满足大小顺序, 这时right先减一, 和left重合
	// 此时right 的值失真,可以和0 交换了
	arr[0], arr[right] = arr[right], arr[0]
	if right > 1 {
		quickSort(arr[:right])
	}
	if right+1 < n-1 {
		quickSort(arr[right+1:])
	}
}

/*
数组中的第K个最大元素
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
*/
func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k < 0 || k > n {
		return -1
	}
	quickSort(nums)
	return nums[n-k]
}
