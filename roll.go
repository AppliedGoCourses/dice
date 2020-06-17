package dice

import (
	"math/rand"
	"time"
)

func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

func RollMulti(num, sides int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += Roll(sides)
	}
	return sum
}

func Seed(n int64) {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
	}
	rand.Seed(n)
}
