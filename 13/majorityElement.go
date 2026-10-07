package main

import "fmt"

func majorityElement1(nums []int) int {
	// var most int
	count_arr := make([]int, len(nums))
	// a := 0
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for j := 0; j < len(nums); j++ {
			if a == nums[j] {
				count_arr[i]++
			}
		}
	}
	return GetMaxFromArr(count_arr)
}

func GetMaxFromArr(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums[len(nums)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main1() {
	arr := []int{1, 2, 2, 2, 2, 2, 3, 3, 3}
	a := majorityElement1(arr)
	fmt.Println(a)
}
