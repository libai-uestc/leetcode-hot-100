package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	minVal := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(num, maxVal*num)
		minVal = min(num, minVal*num)
		ans = max(ans, maxVal)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
