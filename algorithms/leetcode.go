package main

func main() {

}

func twoSum(nums []int, target int) []int {
	var res []int
	for index, i := range nums {
		for jndex, j := range nums {
			if i+j == target {
				res = append(res, index, jndex)
				return res
			}
		}
	}
	return res
}
