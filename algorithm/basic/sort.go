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

// 设定一个基准值P, 通过一次排序使左变的元素都小于P, 右边的元素都大于P
func partition(arr []int, left, right int) int {
	p := arr[right]                 // 取最末元素为基准值
	i := left                       // 用于交换, 保证第i个值小于P
	for j := left; j < right; j++ { // 最有一个元素不需要判断
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i] // 保证第i个元素比P小
			i++
			// arr[i], a
		}
	}
	// 最后一次i累加之后并没有交换到比P小的值, 需要和最后的P位置做一下交换
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
