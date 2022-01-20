package arrays

import "fmt"

//"fmt"

func removeDuplicates(nums []int) int {
	doubles := 0

	length := len(nums)

	if length == 1 {
		return doubles
	}

	for i := 0; i < length-1; i++ {
		fmt.Println("nums[i] ", nums[i])
		fmt.Println("nums[i+1] ", nums[i+1])

		for nums[i] == nums[i+1] {

			if nums[i+1] == 0 {
				break
			}

			for k := i; k < length-1; k++ {
				fmt.Println("k ", k)
				nums[k] = nums[k+1]

				if k+1 == length-1 {
					nums[k+1] = 0
				}
			}
			fmt.Println("nums after k ", nums)

			doubles++
		}
	}

	return doubles
}
