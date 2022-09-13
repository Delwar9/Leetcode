package main

func main() {
	nums := []int{1, 1, 2}
	println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	length, newLength := len(nums), 1
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}
