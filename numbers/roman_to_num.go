package numbers

func romanToInt(s string) int {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0

	reversedString := ""

	for _, run := range s {
		reversedString = string(run) + reversedString
	}

	var temp rune

	for _, run := range reversedString {
		if temp != 0 && nums[run] < nums[temp] {
			num -= nums[run]
		} else {
			num += nums[run]
		}

		temp = run
	}

	return num
}
