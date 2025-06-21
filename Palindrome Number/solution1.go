
func isPalindrome(x int) bool {
	if x < 0  || x == 10 {
		return false
	} else if x >= 0 && x <= 9 {
		return true
	}
	realNum := x
	reverseNum := 0

	for x >= 10 {
		modulus := x % 10
		reverseNum = (reverseNum * 10) + modulus
		x = x / 10
	}

	reverseNum = (reverseNum * 10) + x

	if reverseNum == realNum {
		return true
	} else {
		return false
	}
}