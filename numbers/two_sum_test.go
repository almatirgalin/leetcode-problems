package numbers

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}

	res := twoSum(nums, 6)

	fmt.Println(res)
}
