package test

import (
	"testing"
)

// https://leetcode.cn/problems/longest-consecutive-sequence/submissions/533769642/?envType=study-plan-v2&envId=top-100-liked
func TestLongestSequence(t *testing.T) {
	slice := []int{100, 4, 200, 1, 3, 2, 101, 102, 103, 104, 105}
	consecutive := longestConsecutive(slice)
	println(consecutive)
}

func longestConsecutive(nums []int) int {
	// slice 2map
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for k, _ := range m { //便利m，因为m是去重之后的
		//k就是数组中的一个数字，-1就是这个数的前一个数字，如果不存在，那说明当前数字前面没有了，它就是序列的起点，接下来只用往后找了
		long := 0
		if !m[k-1] { //这一步很关键

			flag := true
			for flag {
				if m[k] {
					long++
					k++
				} else {
					flag = false
				}
				if long > ans {
					ans = long //更新ans 目的是让ans记录当前最长序列的长度
				}
			}

		}
	}
	return ans
}
