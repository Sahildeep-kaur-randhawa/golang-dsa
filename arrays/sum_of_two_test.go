package arrays

import (
	"reflect"
	"testing"
)

// Test cases of sum of two
func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}}, // no solution case
	}

	for _, test := range tests {
		result := TwoSum(test.nums, test.target)

		// Only check values, not the order
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("TwoSum(%v, %d) = %v; expected %v", test.nums, test.target, result, test.expected)
		}
	}
}

// Benchmarking of sum of two numbers
func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}
	target := 19998 // nums[9999] + nums[9999]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
