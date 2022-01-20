package arrays

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := 5

	removeDuplicates(nums)

	fmt.Println(nums)

	expectedNums := []int{1, 2}

	for i := 0; i < k; i++ {
		if expectedNums[i] != nums[i] {
			t.Error("Error")
		}
	}
}
