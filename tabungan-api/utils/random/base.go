package random

import (
	"math/rand"
	"time"
)

const (
	alphabetChars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars    = "0123456789"
	allChars        = alphabetChars + numericChars
	emailExtensions = "com net org biz info edu"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// generateString generates a random string using given characters and length
func generateString(characters string, length int) string {
	var result string

	for i := 0; i < length; i++ {
		result += string(characters[r.Intn(len(characters))])
	}

	return result
}

// GenerateAlphabetString generates a random string using alphabet characters and given length
func GenerateAlphabetString(length int) string {
	return generateString(alphabetChars, length)
}

// GenerateAlphanumericString generates a random string using alphabet and numeric characters and given length
func GenerateAlphanumericString(length int) string {
	return generateString(allChars, length)
}

// GenerateBool generates a random boolean value
func GenerateBool() bool {
	return r.Float32() < 0.5
}

// GenerateFromSet generates a random value from given array
func GenerateFromSet(array []interface{}) interface{} {
	return array[r.Intn(len(array))]
}

// GenerateNumber generates a random number between min and max (inclusive)
func GenerateNumber(min, max int) int {
	return r.Intn(max-min+1) + min
}

// GenerateNumericString generates a random string using numeric characters and given length
func GenerateNumericString(length int) string {
	return generateString(numericChars, length)
}
