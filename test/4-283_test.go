package test

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	ints := []int{0, 1, 0, 3, 12}
	moveZeroes(ints)
	fmt.Println(ints)
}
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 { //1、将所有不是0的值都放到最左边
			nums[j] = nums[i]
			j++ //被放置的位置的索引随着被使用而移动
		}
	}

	for i := j; i < len(nums); i++ { //2、将所有0都放到最右边
		nums[i] = 0
	}
}
