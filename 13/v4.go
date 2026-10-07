package main

import "sort"

func majorityElement3(nums []int) int {
	//获取数组的大小
	n := len(nums)

	//对数组排序
	sort.Ints(nums)

	//返回排序后的中间元素
	return nums[n/2]
}
