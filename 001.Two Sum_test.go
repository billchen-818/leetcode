/**
 * @Author: "billchen-818"
 * @Description:
 * @File: 001.Two Sum_test
 * @Version: 1.0.0
 * @Date: 2020/12/4 14:55
 */

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {
		if n, ok := m[target-value]; ok {
			return []int{n, index}
		}
		m[value] = index
	}

	return []int{}
}

func TestTwoSum(t *testing.T)  {
	a := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	actual := TwoSum(a, target)

	assert.Equal(t, expected, actual)
}
