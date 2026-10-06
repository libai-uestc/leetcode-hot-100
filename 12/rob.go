package main

import "fmt"

// func rob(nums []int) int {
// 	// 1 10 1 1 11 2

// }
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// prev2 表示偷到第 i-2 间房屋时的最高金额
	// prev1 表示偷到第 i-1 间房屋时的最高金额
	prev2, prev1 := 0, 0

	for _, num := range nums {
		// 当前房屋的选择：偷（prev2 + num）或者 不偷（保持 prev1）
		current := max(prev2+num, prev1)

		// 更新记录，为下一个房屋的计算做准备
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

//
//

// 辅助函数：求两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 10, 1, 1, 11, 2}
	a := rob(nums)
	fmt.Println(a)
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := range nums {
		memo[i] = -1
	}
	return rob_rec(nums, memo, len(nums)-1)
}

func rob_rec(nums, memo []int, i int) int {
	if i < 0 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}

	res := max(rob_rec(nums, memo, i-1), rob_rec(nums, memo, i-2)+nums[i])
	memo[i] = res
	return res
}

// func max(a,b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
