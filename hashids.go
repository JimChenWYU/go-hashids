package hashids

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Hashids contains HashidsConfig
type Hashids struct {
	options *HashidsConfig
}

// NewHashidsObjectSimple creates a new default NewHashidsObjectSimple
func NewHashidsObjectSimple() *Hashids {
	return NewHashidsObject(NewDefaultHashidsConfig())
}

// NewHashidsObject creates a new NewHashidsObject
func NewHashidsObject(options *HashidsConfig) *Hashids {
	return &Hashids{
		options: options,
	}
}

// Encode hashes an array of int to a string containing at least MinLength characters taken from the Alphabet.
// Use Decode using the same Alphabet and Salt to get back the array of int.
func (this *Hashids) Encode(numbers []int) string {
	numbersInt64 := make([]int64, 0, len(numbers))
	for _, number := range numbers {
		numbersInt64 = append(numbersInt64, int64(number))
	}

	result, err := this.EncodeInt64(numbersInt64)

	if err != nil {
		panic(err)
	}

	return result
}

// Decode unhashes the string passed to an array of int.
// It is symmetric with Encode if the Alphabet and Salt are the same ones which were used to hash.
// MinLength has no effect on Decode.
func (this *Hashids) Decode(hash string) []int {
	result64, err := this.DecodeInt64(hash)
	if err != nil {
		panic(err)
	}

	result := make([]int, 0, len(result64))
	for _, id := range result64 {
		result = append(result, int(id))
	}

	return result
}

// EncodeHex hashes a hexadecimal string to a string containing at least MinLength characters taken from the Alphabet.
// A hexadecimal string should not contain the 0x prefix.
// Use DecodeHex using the same Alphabet and Salt to get back the hexadecimal string.
//
// Each hex nibble is encoded as an integer in range [16, 31].
func (this *Hashids) EncodeHex(hex string) string {
	result, err := this.EncodeHexWithError(hex)
	if err != nil {
		panic(err)
	}

	return result
}

// DecodeHex unhashes the string passed to a hexadecimal string.
// It is symmetric with EncodeHex if the Alphabet and Salt are the same ones which were used to hash.
//
// Each hex nibble is decoded from an integer in range [16, 31].
func (this *Hashids) DecodeHex(hash string) string {
	result, err := this.DecodeHexWithError(hash)
	if err != nil {
		panic(err)
	}

	return result
}

// see EncodeHex
func (this *Hashids) EncodeHexWithError(hex string) (string, error) {
	for _, b := range hex {
		switch {
		case (b >= '0') && (b <= '9'):
		case (b >= 'a') && (b <= 'f'):
		case (b >= 'A') && (b <= 'F'):
		default:
			return "", errors.New("invalid hex digit")
		}
	}

	hexLength := len(hex)
	hexInt := make([]int, 0, int(math.Ceil(float64(hexLength)/12)))
	for i, step := 0, 12; i < hexLength; i += step {
		rightIndex := i + step
		if rightIndex > hexLength {
			rightIndex = hexLength
		}
		buffer := "1" + hex[i:rightIndex]
		n, err := strconv.ParseUint(buffer, 16, len([]byte(buffer))*4)

		if err != nil {
			return "", err
		}

		hexInt = append(hexInt, int(n))
	}

	return this.Encode(hexInt), nil
}

// see DecodeHex
func (this *Hashids) DecodeHexWithError(hash string) (string, error) {
	resultInt64, err := this.DecodeInt64(hash)
	if err != nil {
		return "", err
	}

	var (
		buffer     string
		bufferRune []rune
		resultRune []rune
	)
	for _, s := range resultInt64 {
		buffer = fmt.Sprintf("%x", s)
		// remove the first character '1'
		bufferRune = stringToRuneArray(buffer[1:])
		resultRune = append(resultRune, bufferRune...)
	}

	return string(resultRune), nil
}

// EncodeInt64 hashes an array of int64 to a string containing at least MinLength characters taken from the Alphabet.
// Use DecodeInt64 using the same Alphabet and Salt to get back the array of int64.
func (this *Hashids) EncodeInt64(numbers []int64) (string, error) {
	if len(numbers) == 0 {
		return "", errors.New("can not encoding empty array of numbers")
	}
	for _, n := range numbers {
		if n < 0 {
			return "", errors.New("negative number not supported")
		}
	}

	var (
		alphabet     = this.options.alphabet
		salt         = this.options.salt
		numbersIdInt = int64(0)
	)
	for i := 0; i < len(numbers); i++ {
		numbersIdInt += numbers[i] % int64(i+100)
	}

	lottery := alphabet[numbersIdInt%int64(len(alphabet))]
	result := []rune{lottery}
	buffer := make([]rune, len(alphabet)+len(salt)+1)
	var last []rune

	for i, number := range numbers {
		buffer = buffer[:1]
		buffer[0] = lottery
		buffer = append(buffer, salt...)
		buffer = append(buffer, alphabet...)
		alphabet = consistentShuffle(alphabet, buffer[:len(alphabet)])
		last = hash(number, alphabet)
		result = append(result, last...)

		if i+1 < len(numbers) {
			number %= int64(last[0]) + int64(i)
			result = append(result, this.options.seps[number%int64(len(this.options.seps))])
		}
	}
	if len(result) < this.options.minLength {
		guardIndex := (numbersIdInt + int64(result[0])) % int64(len(this.options.guards))
		guard := this.options.guards[guardIndex]

		result = append([]rune{guard}, result...)
		if len(result) < this.options.minLength {
			guardIndex = (numbersIdInt + int64(result[2])) % int64(len(this.options.guards))
			guard = this.options.guards[guardIndex]
			result = append(result, guard)
		}
	}

	halfLength := len(alphabet) / 2
	for len(result) < this.options.minLength {
		alphabet = consistentShuffle(alphabet, alphabet)
		result = append(alphabet[halfLength:], result...)
		result = append(result, alphabet[0:halfLength]...)

		excess := len(result) - this.options.minLength
		if excess > 0 {
			begin := excess / 2
			result = result[begin : begin+this.options.minLength]
		}
	}

	return string(result), nil
}

// DecodeInt64 unhashes the string passed to an array of int64.
// It is symmetric with EncodeInt64 if the Alphabet and Salt are the same ones which were used to hash.
// MinLength has no effect on DecodeInt64.
func (this *Hashids) DecodeInt64(hash string) ([]int64, error) {
	var (
		hashes [][]rune
		i      = 0
	)

	hashes = splitRunes(stringToRuneArray(hash), this.options.guards)

	if len(hashes) == 2 || len(hashes) == 3 {
		i = 1
	}

	result := make([]int64, 0, 10)

	hashBreakdown := hashes[i]
	if len(hashBreakdown) > 0 {
		lottery := hashBreakdown[0]
		hashBreakdown = hashBreakdown[1:]
		hashes = splitRunes(hashBreakdown, this.options.seps)
		alphabet := copyRuneSlice(this.options.alphabet)
		buffer := make([]rune, len(alphabet)+len(this.options.salt)+1)
		for _, subHash := range hashes {
			buffer = buffer[:1]
			buffer[0] = lottery
			buffer = append(buffer, this.options.salt...)
			buffer = append(buffer, alphabet...)
			alphabet = consistentShuffle(alphabet, buffer[:len(alphabet)])
			number, err := unhash(subHash, alphabet)
			if err != nil {
				return nil, err
			}
			result = append(result, number)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("unexpected hash string")
	}

	return result, nil
}
