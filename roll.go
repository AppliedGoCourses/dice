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

func RollUntilHighest(sides int) int {
	i, res := 0, 0
	for ; res < sides; i++ {
		res = Roll(sides)
	}
	return i
}

func Seed(n int64) {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
	}
	rand.Seed(n)
}
