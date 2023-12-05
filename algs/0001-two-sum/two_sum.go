package two_sum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))

	for i, n := range nums {
		if pos, ok := hash[target-n]; ok {
			return []int{pos, i}
		}

		hash[n] = i
	}

	return []int{}
}
