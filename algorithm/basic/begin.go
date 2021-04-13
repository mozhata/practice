package basic

import (
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

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
逆波兰表达式求值
https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 取出元素, 入栈, 直到遇到非数, 出栈两个元素, 计算后在入栈
*/
func EvalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var tmp int
			switch tokens[i] {
			case "+":
				tmp = a + b
			case "-":
				tmp = a - b
			case "*":
				tmp = a * b
			case "/":
				tmp = a / b
			}
			stack = append(stack, tmp)
		} else {
			stack = append(stack, val)
		}
	}
	return stack[0]
}

/*
1,2,3    7,4,1
4,5,6 => 8,5,2
7,8,9    9,6,3

1: 对角线对折翻转互换, 2,4 互换 (x,y) => (y,x)
2: 中线对折互换 1,7 互换 (x,y) => (x, n-y-1)
https://leetcode-cn.com/problems/rotate-matrix-lcci/

*/
func Rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	// 对角线互换
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 中线互换
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			matrix[j][i], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[j][i]
		}
	}
}

/*
字符串解码
输入：s = "3[a]2[bc]"
输出："aaabcbc"
https://leetcode-cn.com/problems/decode-string/
*/
func DecodeStr(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if s[i] != ']' {
			stack = append(stack, s[i])
			continue
		}
		// 是 ']', 把"[]"内的字符找出来,
		tmp := make([]byte, 0)
		var b byte
		for {
			// pop 掉 '['
			b = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if b == '[' {
				break
			}
			tmp = append(tmp, b)
		}
		// 找到数字部分
		var idx = 1
		for len(stack) >= idx {
			if stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
				continue
			}
			break
		}
		nums := stack[len(stack)-idx+1:]
		stack = stack[:len(stack)-idx+1]
		num, _ := strconv.Atoi(string(nums))
		// 完成翻译
		// 压辉栈内
		for ii := 0; ii < num; ii++ {
			for j := 0; len(tmp)-j > 0; j++ {
				stack = append(stack, tmp[len(tmp)-j-1])
			}
		}
	}
	return string(stack)
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

/*
三角形最小路径和
https://leetcode-cn.com/problems/triangle/
*/
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	// f[i][j] 表示从顶部走到(i,j)位置的最小路径和
	f := make([][]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		f[i] = make([]int, i+1)
	}
	f[0][0] = triangle[0][0]
	// 计算
	for i := 1; i < n; i++ {
		// 左边(j=0)
		f[i][0] = f[i-1][0] + triangle[i][0]
		// 右边(j=i)
		f[i][i] = f[i-1][i-1] + triangle[i][i]
		// 中间部分
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
		}
	}

	// 在f[n-1][0]到f[n-1][n-1] 中找最小值
	ans := f[n-1][0]
	for i := 1; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

/*
不同路径
https://leetcode-cn.com/problems/unique-paths/
*/
func UniquePaths(m int, n int) int {
	if m < 2 || n < 2 {
		return 1
	}
	// f[i][j] 表示走到(i,j)的路径数, f[i][j] = f[i-1,j] + f[i][j-1]
	// 初始化
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		f[i][0] = 1
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i][j-1] + f[i-1][j]
		}
	}
	return f[m-1][n-1]
}

/*
最小覆盖子串
https://leetcode-cn.com/problems/minimum-window-substring/
*/
// 超出时间限制了, 下次再来
func MinWindowV1(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var (
		mt         = make(map[byte]int)
		lenT, lenS int

		bestL, bestR int // 最小窗口的左右indx
		start, end   int // 临时窗口的左右indx
		matchNum     int
	)
	for i := range t {
		mt[t[i]] += 1
		lenT++
	}

	lenS = len(s)
	bestR = lenS // 初始化为一个较大值
	if lenS < lenT {
		return ""
	}
	shouldMatchNum := len(mt)
	min := lenS + 1
	win := make(map[byte]int)
	for end < lenS {
		c := s[end]
		end++
		if mt[c] > 0 {
			win[c] += 1
			if win[c] == mt[c] {
				matchNum++
			}
		}
		// win 里面内满的时候
		for matchNum == shouldMatchNum {
			if end-start < min {
				min = end - start
				bestL, bestR = start, end
			}
			c := s[start]
			start++
			if win[c] > 0 {
				if win[c] == mt[c] {
					matchNum--
				}
				win[c]--
			}
		}

	}

	if min > lenS {
		return ""
	}
	return s[bestL:bestR]
}

/*
无重复字符的最长子串
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
func LengthOfLongestSubstring(s string) int {
	lenS := len(s)
	if lenS == 0 {
		return 0
	}
	var (
		left, right int
		win         = make(map[byte]int)
		maxLen      int
	)
	for right < lenS {
		c := s[right]
		right++
		win[c]++
		for win[c] > 1 {
			d := s[left]
			win[d]--
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}
