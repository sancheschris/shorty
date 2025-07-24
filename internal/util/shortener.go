package util

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortCode() string {
	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""

	for i := 0; i < 6; i++ {
		result += string(base62[rng.Intn(len(base62))])
	}

	return result
}
