package test_common

import (
	"fmt"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	nums := []int{1, 2, 3}
	permute(nums)
	fmt.Println(result_)
}

// 1 <= nums.length <= 6 // -10 <= nums[i] <= 10  // nums 中的所有整数 互不相同

func permute(nums []int) [][]int {
	track := make([]int, 0, 0)
	result_ = make([][]int, 0, 0) //为了每次清空上一次的结果
	backtrace(nums, track)
	return result_
}

var result_ [][]int

func backtrace(nums []int, track []int) {
	track_ := make([]int, len(track))
	copy(track_, track)
	result_ = append(result_, track_) //记录组合

	for _, num := range nums {
		//避免有重复的数字
		if has(track, num) {
			continue
		}
		track = append(track, num) //记录路径
		backtrace(nums, track)
		track = track[0 : len(track)-1]
	}
}

func has(track []int, currrnt int) bool {
	for _, value := range track {
		if value == currrnt {
			return true
		}
	}
	return false
}
