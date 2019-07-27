package hashids

import "strings"

func uniqueCharacter(alphabet string) string {
	var result strings.Builder
	for i, alphabetLen := 0, len(alphabet); i < alphabetLen; i++ {
		if strings.IndexByte(result.String(), byte(alphabet[i])) == -1 {
			result.WriteByte(byte(alphabet[i]))
		}
	}
	return result.String()
}

func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func consistentShuffle(alphabet []rune, salt []rune) []rune {
	var (
		integer int
		j       int
	)

	if len(salt) == 0 {
		return alphabet
	}

	result := copyRuneSlice(alphabet)
	for i, v, p := len(result)-1, 0, 0; i > 0; i, v = i-1, v+1 {
		v %= len(salt)
		integer = int(salt[v])
		p += integer
		j = (integer + v + p) % i

		result[i], result[j] = result[j], result[i]
	}

	return result
}

func copyRuneSlice(r []rune) []rune {
	result := make([]rune, len(r))
	copy(result, r)
	return result
}

func stringToRuneArray(str string) []rune {
	return []rune(str)
}
