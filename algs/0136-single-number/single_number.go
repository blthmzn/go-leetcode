package single_number

type Hash map[int]struct{}

func singleNumber(nums []int) int {
	hashSize := (len(nums) / 2) + 1
	hash := make(Hash, hashSize)

	for _, v := range nums {
		if _, ok := hash[v]; ok {
			delete(hash, v)
		} else {
			hash[v] = struct{}{}
		}
	}

	var result int
	for k := range hash {
		result = k
	}

	return result
}
