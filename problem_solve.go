package main

import "fmt"

// twoSum returns indices of two numbers in the given slice that add up to the target number.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
