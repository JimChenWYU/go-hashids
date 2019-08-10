package hashids

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Unique a string character
func uniqueCharacter(alphabet string) string {
	var result strings.Builder
	for i, alphabetLen := 0, len(alphabet); i < alphabetLen; i++ {
		if strings.IndexByte(result.String(), byte(alphabet[i])) == -1 {
			result.WriteByte(byte(alphabet[i]))
		}
	}
	return result.String()
}

// Remove a string spaces
func removeSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}

// Shuffle a string
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

// Hash func
func hash(input int64, alphabet []rune) []rune {
	result := make([]rune, 0)
	for {
		result = append(result, alphabet[input%int64(len(alphabet))])
		input = input / int64(len(alphabet))
		if input == 0 {
			break
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Unhash func
func unhash(input, alphabet []rune) (int64, error) {
	result := int64(0)
	for _, inputRune := range input {
		alphabetPos := strings.IndexRune(string(alphabet), inputRune)
		if alphabetPos == -1 {
			return 0, errors.New("alphabet used for hash was different")
		}
		result = result*int64(len(alphabet)) + int64(alphabetPos)
	}
	return result, nil
}

// Split a string
func splitRunes(input, seps []rune) [][]rune {
	splitPoses := make([]int, 0)
	for pos, inputRune := range input {
		for _, r := range seps {
			if inputRune == r {
				splitPoses = append(splitPoses, pos)
				break
			}
		}
	}

	result := make([][]rune, 0, len(splitPoses)+1)
	inputLeft := input[:]
	for _, splitIndex := range splitPoses {
		splitIndex -= len(input) - len(inputLeft)
		result = append(result, inputLeft[:splitIndex])
		inputLeft = inputLeft[splitIndex+1:]
	}
	result = append(result, inputLeft)

	return result
}

// Copy Slice
func copyRuneSlice(r []rune) []rune {
	result := make([]rune, len(r))
	copy(result, r)
	return result
}

// Transfer a string to rune array
func stringToRuneArray(str string) []rune {
	return []rune(str)
}

// Debug func
func debug(input ...interface{}) {
	fmt.Println(input...)
	os.Exit(0)
}
