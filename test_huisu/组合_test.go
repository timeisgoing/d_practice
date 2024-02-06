package huisu_test

import (
	"fmt"
	"testing"
)

// 回溯-2  77. 组合  https://leetcode.cn/problems/combinations/description/
func TestName(t *testing.T) {
	i := combine(1, 1)
	fmt.Println(i)
}

var result_ [][]int

// 意思：在值的范围是[1,n]之间找到k个组合
func combine(n int, k int) [][]int {
	track := make([]int, 0, 0)
	//result_ := make([][]int, 0, 0) // fix 初始化
	backtrace(n, k, 1, track) //1 是因为值的范围是[1,n]
	return result_
}

func backtrace(n, k, start int, track []int) {
	//	从1到n，取k个
	if len(track) == k {
		trackCopy := make([]int, len(track))
		copy(trackCopy, track)
		result_ = append(result_, trackCopy)
	}
	for i := start; i <= n; i++ { //start=1 是因为值的范围是[1,n]
		track = append(track, i)        //add
		backtrace(n, k, i+1, track)     //递归，i+1是因为i已经包含了，不能重复
		track = track[0 : len(track)-1] //回溯
	}
}
