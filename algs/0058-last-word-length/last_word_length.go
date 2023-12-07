package last_word_length

func lengthOfLastWord(s string) int {
	var res int

	i := len(s) - 1

	for range s {
		if string(s[i]) != " " {
			for i > 0 || string(s[i]) != " " {
				res++
				i--
			}

			return res
		}

		i--
	}

	return res
}
