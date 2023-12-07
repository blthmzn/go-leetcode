package number_of_bits

func hammingWeight(num uint32) int {
	var res int

	for num != 0 {
		if num&0x01 == 0x01 {
			res++
		}

		num = num >> 1
	}

	return res
}
