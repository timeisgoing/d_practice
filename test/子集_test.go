package test

import (
	"fmt"
	"testing"
)

// 78. 子集 https://leetcode.cn/problems/subsets/description/
// https://leetcode.cn/problems/subsets/solutions/586887/golangban-ben-hui-su-mo-ban-li-jie-ti-xi-upx4/ 这个题解和我思路一样
func TestName1(t *testing.T) {
	nums := []int{1, 2, 3}
	result1 = make([][]int, 0, 0)
	subsets(nums)
	fmt.Println(result1)
}

var result1 [][]int

func subsets(nums []int) [][]int {
	track := make([]int, 0, 0)
	backtrace1(nums, track, 0)
	return result1
}

func backtrace1(nums, track []int, index int) {
	//track_ := make([]int, 0, 0) // bug
	track_ := make([]int, len(track))
	copy(track_, track)
	result1 = append(result1, track_) //每次都记录一下当前的路径

	//for i := 0; i < len(nums); i++ {// bug
	for i := index; i < len(nums); i++ {
		track = append(track, nums[i]) //记录当前的路径
		backtrace1(nums, track, i+1)   //i+1 排除当前i应算在内了
		//回溯,删除最后那一个
		track = track[0 : len(track)-1]
	}

}
