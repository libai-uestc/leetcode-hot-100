package main

func majorityElement4(nums []int) int {
	res := 0
	//最大频率
	maxCount := 0

	// 使用 make 初始化哈希表，替代 C++ 中的 unordered_map
	num2count := make(map[int]int)

	for _, num := range nums {
		num2count[num]++

		//当前元素的频率大于最大频率
		if num2count[num] > maxCount {
			maxCount = num2count[num]
			res = num
		}
	}

	return res
}
