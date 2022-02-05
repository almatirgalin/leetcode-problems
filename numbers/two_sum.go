package numbers

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, val := range nums {
		if num, ok := m[target-val]; ok {
			return []int{i, num}
		}

		m[val] = i
	}

	return []int{}
}
