package helper

import (
	"math/rand"
)

// XorShift Helper function
func XorShift(r uint64) uint64 {
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	return r
}


const charset = "abcdefghijklmnopqrstuvwxyz"

// Random char generator
func GenerateRandomString(strLength int) string {
	result := make([]byte, strLength)
	for i := 0; i < strLength; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}



func CropString(s string, size int) string {
	if len(s) <= size {
		return s
	}
	return s[:size] + "..."
}
