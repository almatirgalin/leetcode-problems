package numbers

import "testing"

func TestIsPalindrom(t *testing.T) {
	num := 12321

	if !isPalindrome(num) {
		t.Error("12321 is palindrome number")
	}

	num = 1221

	if !isPalindrome(num) {
		t.Error("1221 is palindrome number")
	}

	num = 1242

	if isPalindrome(num) {
		t.Error("1242 isn't palindrome number")
	}
}
