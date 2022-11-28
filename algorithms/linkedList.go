package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 1))
}

func binarySearch(a []int, search int) (result int) {
	var arr []int
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
		arr = append(arr, a...)
	}
	fmt.Println(arr)
	return
}

func searchRange(nums []int, target int) []int {
	return []int{}
}
