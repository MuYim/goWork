package main

import "strconv"

func main() {
	println("Hello, world.")
	println(isPalindrome(10))
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}
