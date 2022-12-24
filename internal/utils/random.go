package utils

import (
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateInLine(min, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Intn(max-min)
}

func GenerateRandomString(size int) string {
	rand.Seed(time.Now().Unix())

	var str = make([]rune, size)
	for i := 0; i < size; i++ {
		str = append(str, runes[rand.Intn(len(runes)-1)])
	}

	return string(str)
}
