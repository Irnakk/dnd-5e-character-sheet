package dice

import (
	"errors"
	"math/rand"
)

var (
	dice = map[int]struct{}{
		4:   {},
		6:   {},
		8:   {},
		10:  {},
		12:  {},
		20:  {},
		100: {},
	}
)

func Dice(number, value int) (int, error) {
	if _, found := dice[value]; !found {
		return 0, errors.New("no such die")
	}

	if number < 1 {
		return 0, errors.New("invalid number of rolls")
	}

	sum := 0

	for i := 0; i < number; i++ {
		die := rand.Intn(value) + 1 // Since it was [0; value), it'll be [1, value]
		sum += die
	}

	return sum, nil
}
