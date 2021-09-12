package dice

import (
	"encoding/json"
	"fmt"
)

type RollInfo struct {
	Number int
	Value  int
	Result int
}

type StatsInfo struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
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

func RollStats() (StatsInfo, error) {
	resultStats := StatsInfo{}

	roll, err := DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Strength = roll

	roll, err = DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Dexterity = roll

	roll, err = DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Constitution = roll

	roll, err = DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Intelligence = roll

	roll, err = DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Wisdom = roll

	roll, err = DiceRoll(4, 6)
	if err != nil {
		fmt.Printf("Error in rolling 4d6:\t%v\n", err)
		return StatsInfo{}, err
	}
	resultStats.Charisma = roll

	return resultStats, nil
}
