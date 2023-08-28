package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	breakPoint := -1

	// Find the index where the sequence goes from non-increasing to increasing
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			breakPoint = i
			break
		}
	}

	if breakPoint == -1 {
		reverse(nums, 0, n-1)
		return
	}

	// Find the next greater element to the right of breakPoint
	for i := n - 1; i > breakPoint; i-- {
		if nums[i] > nums[breakPoint] {
			nums[i], nums[breakPoint] = nums[breakPoint], nums[i]
			break
		}
	}

	// Reverse the elements to the right of breakPoint
	reverse(nums, breakPoint+1, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
