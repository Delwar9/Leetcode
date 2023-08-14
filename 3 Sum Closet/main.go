package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[n-1]

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1

		for j < k {
			currentSum := nums[i] + nums[j] + nums[k]

			if currentSum <= target {
				j++
			} else {
				k--
			}

			if abs(closest-target) > abs(currentSum-target) {
				closest = currentSum
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := threeSumClosest(nums, target)
	fmt.Println(result)
}
