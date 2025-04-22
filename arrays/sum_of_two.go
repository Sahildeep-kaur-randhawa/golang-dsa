package arrays

// Finding sum of two numbers equals to the target number in array
// Time complexity is : O(n)
func TwoSum(nums []int, target int) []int {
	output := make([]int, 0, 2)
	complement := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, val := complement[num]; val {
			output = append(output, complement[num], i)
		}
		complement[target-num] = i
	}
	return output
}
