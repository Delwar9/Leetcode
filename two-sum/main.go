package main

import (
	"fmt"
	"testing"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	Test_Problem1(&testing.T{})
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

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
