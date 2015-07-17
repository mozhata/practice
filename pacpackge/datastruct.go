package pacpackge

import (
	"fmt"
	"time"
)

var slice = []int{2, 3, -4, 2, 3, 1}

func MaxSubsequence3() {
	thisSum, MaxSum := 0, 0
	for i := 0; i < len(slice); i++ {
		for j := i; j < len(slice); j++ {
			thisSum = 0
			for k := i; k <= j; k++ {
				thisSum += slice[k]
				if thisSum > MaxSum {
					MaxSum = thisSum
				}
			}
		}
	}
}
func MaxSubsequence2() {
	thisSum, MaxSum := 0, 0
	for i := 0; i < len(slice); i++ {
		thisSum = 0
		for j := i; j < len(slice); j++ {
			thisSum += slice[j]
			if thisSum > MaxSum {
				MaxSum = thisSum
			}
		}
	}
}
func MaxSubsequenceN() {
	thisSum, MaxSum := 0, 0
	for i := 0; i < len(slice); i++ {
		thisSum += slice[i]
		if thisSum > MaxSum {
			MaxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
}
func Ttime() {
	fmt.Println("test N3...")
	begin := time.Now()
	for i := 0; i < 10000000; i++ {
		MaxSubsequence3()
	}
	end := time.Now()
	fmt.Println(end.Sub(begin), "\ntest N2...")
	begin = time.Now()
	for i := 0; i < 10000000; i++ {
		MaxSubsequence3()
	}
	end = time.Now()
	fmt.Println(end.Sub(begin), "\ntest N...")
	begin = time.Now()
	for i := 0; i < 10000000; i++ {
		MaxSubsequence3()
	}
	end = time.Now()
	fmt.Println(end.Sub(begin))
}

// func TrySum() {
// 	numStr := "123456789"
// 	for calNum := 1; calNum < 9; calNum++ {
// 		for index := 0; index < 9; index++ {

// 		}
// 	}
// }
