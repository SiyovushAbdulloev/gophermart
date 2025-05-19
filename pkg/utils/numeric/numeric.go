package numeric

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func IsNumberic(value string) bool {
	result := false
	for _, v := range []rune(value) {
		for _, n := range numbers {
			if string(v) == n {
				result = true
			}
		}
	}

	return result
}
