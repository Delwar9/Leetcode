package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	dfs(nums, 0, []int{}, &ans)
	return ans
}

func dfs(nums []int, s int, path []int, ans *[][]int) {
	*ans = append(*ans, append([]int(nil), path...))

	for i := s; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1, path, ans)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}
