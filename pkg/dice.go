package dice

import (
	"errors"
	"math/rand"
)

var (
	dice_types = map[int]struct{}{
		4:   {},
		6:   {},
		8:   {},
		10:  {},
		12:  {},
		20:  {},
		100: {},
	}
)

func dice(number, value int) (int, error) {
	if _, found := dice_types[value]; !found {
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

func diceVerbose(number, value int) ([]int, error) {
	if _, found := dice_types[value]; !found {
		return []int{}, errors.New("no such die")
	}

	if number < 1 {
		return []int{}, errors.New("invalid number of rolls")
	}

	result := make([]int, number)

	for i := 0; i < number; i++ {
		die, err := dice(1, value)

		if err != nil {
			return []int{}, errors.New("error in rolling a single die")
		}

		result = append(result, die)
	}

	return result, nil
}
