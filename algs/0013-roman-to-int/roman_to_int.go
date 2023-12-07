package roman_to_int

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// input: "MCMXCIV"
// output: 1994
func romanToInt(s string) (result int) {
	var prev int

	for i := len(s) - 1; i >= 0; i-- {
		curr := romans[string(s[i])]
		if curr >= prev {
			result += curr
		} else {
			result -= curr
		}

		prev = curr
	}

	return
}
