package main

import "fmt"

func majorityElement11(nums []int) int {
	res := 0
	maxCount := 0
	num2count := make(map[int]int)
	for i, num := range nums {
		num2count[num]++
		fmt.Println("i = ", i, "num2count[num] = ", num2count[num])
		if num2count[num] > maxCount {
			maxCount = num2count[num]
			res = num
		}
	}
	fmt.Println("num2count[1] = ", num2count[1])
	fmt.Println("num2count[2] = ", num2count[2])
	fmt.Println(num2count)
	return res

}
func main() {
	arr := []int{1, 2, 2, 2, 2, 2, 3, 3, 3}
	a := majorityElement(arr)
	fmt.Println(a)
}
