package leetcode

import "math"

func minOperations(nums []int) int {
	result := 0

	// Make a map of unique number frequency
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	// Check if each number in the freq map is divisible by 3 and/or 2
	for _, count := range freq {

		if( count == 1 ) {
			return -1
		}

		result += int(math.Ceil(float64(count) / 3.0))

		// // Check if whole count divisible by 3
		// if count % 3 == 0 {
		// 	result += count / 3
		// 	continue
		// }

		// // Check if partial count divisible by 3 then 2
		// if (count%3) % 2 == 0 {
		// 	result += count / 3
		// 	result += (count%3) / 2
		// 	continue
		// }

		// // Off by one case
		// if ((count - 2) % 3) % 2 == 0 {
		// 	result ++
		// 	result += (count - 2) / 3
		// 	result += ((count - 2) % 3) / 2
		// 	continue
		// }

		// // Check if divisible by 2
		// if count % 2 == 0 {
		// 	result += count / 2
		// 	continue
		// }

		// return -1

	}

  return result
}