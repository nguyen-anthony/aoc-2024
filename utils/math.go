package utils

func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

