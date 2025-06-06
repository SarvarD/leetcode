// https://leetcode.com/problems/two-sum/description/
package leetcode

import (
	"fmt"
	"log"
)

func twoSum(nums []int, target int) []int {
	tracker := make(map[int]int)
	for idx, num := range nums {
		existingIndexOfComplement, ok := tracker[num]
		if ok {
			return []int{existingIndexOfComplement, idx}
		}

		tracker[target-num] = idx
	}

	// Problem specifies there will always be a solution, so we should never get here
	log.Fatal("Unexpected: Input has no solution")
	return nil
}

func main() {
	res := twoSum([]int{2, 5, 3, 7, 11, 15}, 9)
	fmt.Println(res)
}
