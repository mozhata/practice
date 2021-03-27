package basic

/*
28. 实现 strStr()
https://leetcode-cn.com/problems/implement-strstr/
*/
func IndexOfSubStr(haystack, needle string) int {
	subLen := len(needle)
	if subLen == 0 {
		return 0
	}
	haysLen := len(haystack)
	var (
		i, j int
	)
	for i = 0; i < haysLen-subLen+1; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		for j = 0; j < subLen; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == subLen {
			return i
		}
	}
	return -1
}

func TryForI(n int) int {
	var i int
	for i = 0; i < n; i++ {
	}
	return i
}

/*
subsets: <M>
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
*/
// 递归法:
// TODO

// 回溯法
// TODO: 还没看懂

// 二进制法
// TODO
