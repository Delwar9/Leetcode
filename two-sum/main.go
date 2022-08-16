package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {

		if v, found := m[target-num]; found {
			return []int{v, i}
		}

		m[num] = i
	}
	return nil
}
