package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(a, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {
		if n, ok := m[target-value]; ok {
			return []int{n, index}
		}
		m[value] = index
	}

	return []int{}
}
