package numbers

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var reversedNum int

	temp := x

	for temp > 0 {
		reversedNum = reversedNum*10 + temp%10
		temp = temp / 10
	}

	return reversedNum == x
}
