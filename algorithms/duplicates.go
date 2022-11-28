package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		}
	}
	return len(nums) - count
}

func rotate(nums []int, k int) []int {
	return append(nums[len(nums)-k:len(nums)], nums[0:len(nums)-k]...)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(nums, 3))
	fmt.Println(reverseArray(nums))
}

func reverseArray(a []int) []int {
	var temp []int
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(i)
		temp = append(temp, a[i])
	}
	return temp
}
