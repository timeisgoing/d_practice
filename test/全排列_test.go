package test

import (
	"fmt"
	"testing"
)

//回溯-1  LCR 083. 全排列   https://leetcode.cn/problems/VvJkup/solutions/2591310/jing-dian-hui-su-by-song-jia-liang-fzqj/

func TestTimeConsuming(t *testing.T) {
	nums := []int{1, 2, 3}
	result1 := permute(nums)
	fmt.Println(result1)
}

// 1 <= nums.length <= 6 // -10 <= nums[i] <= 10  // nums 中的所有整数 互不相同

var result [][]int

func permute(nums []int) [][]int {
	track := make([]int, 0, 0)
	result = make([][]int, 0, 0) //为了每次清空上一次的结果
	backTrack(nums, track)
	return result
}

func backTrack(nums []int, track []int) {
	// list 暂存全排列的数组
	if len(track) == len(nums) { //长度一样了，那就记录结果
		//result = append(result, track) //有问题，因为track是一个切片，后续的回溯过程会修改track，导致之前添加到result中的全排列也会被修改。这样会导致最终的结果不正确。
		trackCopy := make([]int, len(track))
		copy(trackCopy, track)
		result = append(result, trackCopy)
	}
	for _, num := range nums {
		//避免有重复的数字
		if has(track, num) {
			continue
		}
		track = append(track, num) //添加上路径节点
		backTrack(nums, track)     //再往下递归
		//del，这是回溯算法的关键，回溯时删除刚刚添加的元素
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

func Test2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0 : len(nums)-1])
}
