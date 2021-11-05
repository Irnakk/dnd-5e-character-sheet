/*
Package dice provides the basic structures and functions that allow
to perform dice rolls.
*/
package dice

import (
	"errors"
	"math/rand"
)

var (
	// Starts with a lower-case letter because this struct is not exported
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

// DiceRoll returns a number that represents a
// summary result of a random roll of
// <number> dice that have <value> sides.
func DiceRoll(number, value int) (int, error) {
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

// DiceRollVerbose rolls <value>-sided dice
// <number> times and returns a slice of roll
// results on each iteration.
func DiceRollVerbose(number, value int) ([]int, error) {
	if _, found := dice_types[value]; !found {
		return []int{}, errors.New("no such die")
	}

	if number < 1 {
		return []int{}, errors.New("invalid number of rolls")
	}

	result := make([]int, 0, number)

	for i := 0; i < number; i++ {
		die, err := DiceRoll(1, value)

		if err != nil {
			return []int{}, errors.New("error in rolling a single die")
		}

		result = append(result, die)
	}

	return result, nil
}
