package utils

import (
	"math/rand"
	"time"
)

func GenerateInLine(min, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Intn(max-min)
}
