package dice

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"dnd-5e-character-sheet/src/csdata"
)

type RollInfo struct {
	Number int
	Value  int
	Result int
}

func MarshalRollSingle(number, value int) ([]byte, error) {
	roll_result, err := DiceRoll(number, value)
	if err != nil {
		fmt.Printf("Error in rolling dice:\t%v\n", err)
		return []byte{}, err
	}

	result_obj := RollInfo{
		Number: number,
		Value:  value,
		Result: roll_result,
	}

	marshaledResult, err := json.MarshalIndent(result_obj, "", "	")
	if err != nil {
		fmt.Printf("Error in marshalling response_object:\t%v\n", err)
		return []byte{}, err
	}

	return marshaledResult, nil
}

func RollStats() (csdata.SixStats, error) {
	rolls := make([]int, 6)

	for i := 0; i < 6; i++ {
		roll, err := DiceRoll(4, 6)
		if err != nil {
			fmt.Printf("Error in rolling 4d6 on iteration %v:\t%v\n", i, err)
			return csdata.SixStats{}, err
		}

		rolls[i] = roll
	}

	resultStats := csdata.SixStats{
		Strength:     rolls[0],
		Dexterity:    rolls[1],
		Constitution: rolls[2],
		Intelligence: rolls[3],
		Wisdom:       rolls[4],
		Charisma:     rolls[5],
	}

	return resultStats, nil
}

func Modifier(stat int) (int, error) {
	if stat < 1 {
		return 0, errors.New("stat cannot be less than 1")
	}

	return int(math.Floor((float64(stat) - 10.) / 2.)), nil
}
