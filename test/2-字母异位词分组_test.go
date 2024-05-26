package test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	str := "abba"
	split := strings.Split(str, "")
	fmt.Printf("%v\n", split) //[a b b a]
	sort.Strings(split)
	fmt.Printf("%v\n", split) //[a a b b]
	println("========================")
	//sort.strings()的不足，它是区分大小写的
	str2 := "afBA"
	split2 := strings.Split(str2, "")
	fmt.Printf("%v\n", split2) //[a b b a]
	sort.Strings(split2)
	fmt.Printf("%v", split2) //[a a b b]

	println("========================")

	//如果需要进行不区分大小写的字符串排序，可以使用strings.ToLower()函数将字符转换为小写进行比较。
	arr := []string{"apple", "Banana", "cherry", "dATE", "Eggplant", "Fig"}

	sort.Slice(arr, func(i, j int) bool {
		return strings.ToLower(arr[i]) > strings.ToLower(arr[j])
	})

	fmt.Printf("%v", arr)

}

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func TestGroupAnagrams2(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Printf("result:%v \n", result)
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, 0)
	dict := make(map[string][]string)
	for _, str := range strs {
		//将这个字符串排序：方法1，字符串转为一个字节数组进行排序 方法2，字符串切割转为一个字符串数组
		str_slice := strings.Split(str, "")

		fmt.Printf("字符串:%v \n", str_slice)
		sort.Strings(str_slice)
		fmt.Printf("sort:%v \n", str_slice)
		key := slice2str(str_slice)
		dict[key] = append(dict[key], str)
	}
	for _, v := range dict {
		result = append(result, v)
	}
	return result
}

func slice2str(list []string) string {
	result := ""
	for _, v := range list {
		result += v
	}
	return result
}

// 方法2，字符串切割转为一个字符串数组
func groupAnagrams2(strs []string) [][]string {
	result := make([][]string, 0, 0)
	dict := make(map[string][]string)
	for _, str := range strs {
		//将这个字符串排序  :字符串转为一个字节数组进行字节排序，然后再转为一个字符串

		fmt.Printf("字符串:%v \n", str)

		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		fmt.Printf("排序后的字符串:%v \n", string(s))

		dict[string(s)] = append(dict[string(s)], str)
	}
	for _, v := range dict {
		result = append(result, v)
	}
	return result
}
