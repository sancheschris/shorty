package util

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortCode(length int) string {
	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""

	for i := 0; i < length; i++ {
		result += string(base62[rng.Intn(len(base62))])
	}

	return result
}
