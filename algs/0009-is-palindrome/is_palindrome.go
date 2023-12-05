package is_palindrome

func isPalindrome(x int) bool {
	var r int

	for t := x; t > 0; t /= 10 {
		r = r*10 + t%10
	}

	return x == r
}
