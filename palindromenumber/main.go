package main
import "fmt"
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var rev int
	dividedx := x

	for dividedx != 0 {
		//pop
		pop := dividedx % 10
		dividedx /= 10
		//push
		rev = rev*10 + pop
	}

	return x == rev
}

func main() {
	fmt.Println(isPalindrome(121))
}
