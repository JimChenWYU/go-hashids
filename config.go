package hashids

import (
	"fmt"
	"math"
)

const (
	DEFAULT_ALPHABET   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	DEFAULT_SALT       string = ""
	DEFAULT_MIN_LENGTH int    = 0

	SEPS                string  = "cfhistuCFHISTU"
	MIN_ALPHABET_LENGTH int     = 16
	SEP_DIV             float64 = 3.5
	GUARD_DIV           float64 = 3.5
)

type HashidsConfig struct {
	alphabet  []rune
	seps      []rune
	guards    []rune
	salt      []rune
	minLength int

	originalSalt     string
	originalAlphabet string
}

func NewDefaultHashidsConfig() *HashidsConfig {
	return NewHashidConfig(DEFAULT_SALT, DEFAULT_MIN_LENGTH, DEFAULT_ALPHABET)
}

func NewHashidConfig(_salt string, _minLength int, _alphabet string) *HashidsConfig {
	config := &HashidsConfig{
		alphabet:         nil,
		seps:             nil,
		guards:           nil,
		salt:             nil,
		originalSalt:     "",
		originalAlphabet: "",
	}
	config.initialize(_salt, _minLength, _alphabet)
	return config
}

func (this *HashidsConfig) notify(_salt string, _minLength int, _alphabet string) {
	this.initialize(_salt, _minLength, _alphabet)
}

func (this *HashidsConfig) initialize(_salt string, _minLength int, _alphabet string) {
	filterAlphabet := uniqueCharacter(_alphabet)
	filterAlphabet = removeSpaces(_alphabet)

	if len(filterAlphabet) < MIN_ALPHABET_LENGTH {
		panic(fmt.Sprintf("alphabet must contain at least %d different characters", MIN_ALPHABET_LENGTH))
	}

	alphabet := stringToRuneArray(filterAlphabet)
	salt := stringToRuneArray(_salt)
	seps := stringToRuneArray(SEPS)
	var guards []rune

	/*
		`this.seps` should contain only characters present in `this.alphabet`
		`this.alphabet` should not contains `this.seps`
	*/
	for i := 0; i < len(seps); i++ {
		found := -1
		for j, a := range alphabet {
			if a == seps[i] {
				found = j
			}
		}

		if found == -1 {
			seps = append(seps[:i], seps[i+1:]...)
		} else {
			alphabet = append(alphabet[:found], alphabet[found+1:]...)
		}
	}

	seps = consistentShuffle(seps, salt)

	if len(seps) > 0 || float64(len(alphabet))/float64(len(seps)) > SEP_DIV {
		sepsLength := int(math.Ceil(float64(len(alphabet)) / SEP_DIV))
		if sepsLength > len(seps) {
			diff := sepsLength - len(seps)
			seps = append(seps, alphabet[:int(diff)]...)
			alphabet = alphabet[diff:]
		}
	}

	alphabet = consistentShuffle(alphabet, salt)
	guardCount := int(math.Ceil(float64(len(alphabet)) / GUARD_DIV))

	if len(alphabet) > 3 {
		guards = seps[:guardCount]
		seps = seps[guardCount:]
	} else {
		guards = alphabet[:guardCount]
		alphabet = alphabet[guardCount:]
	}

	this.originalAlphabet = _alphabet
	this.originalSalt = _salt

	this.minLength = _minLength
	this.alphabet = alphabet
	this.salt = salt
	this.seps = seps
	this.guards = guards
}

func (this *HashidsConfig) SetMinLength(_minLength int) {
	this.minLength = _minLength
}

func (this *HashidsConfig) SetSalt(_salt string) {
	this.notify(_salt, this.minLength, this.originalAlphabet)
}

func (this *HashidsConfig) SetAlphabet(_alphabet string) {
	this.notify(this.originalSalt, this.minLength, _alphabet)
}
