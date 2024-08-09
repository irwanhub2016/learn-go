// nums := []int{2, 7, 11, 15}
// target := 9
// result := findTwoSum(nums, target)
// fmt.Println(result) // Output: [0 1]

package main

import "fmt"

func main() {
	nums := []int{3, 5, 2, 6, 7, 8}
	target := 9
	result := findTwoSum(nums, target)
	fmt.Println(result)
}

func findTwoSum(nums []int, target int) []int {
	s := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				s = append(s, j)
				s = append(s, i)
			}
		}
	}

	return s
}
