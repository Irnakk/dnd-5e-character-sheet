package dice

import "testing"

func TestDiceRoll(t *testing.T) {

	// test for invalid number of rolls
	result_inv_num, err_inv_num := DiceRoll(-1, 8)

	if result_inv_num != 0 || err_inv_num == nil {
		t.Errorf("DiceRoll(-1, 8) failed; expected %v, got %v", "error", result_inv_num)
	} else {
		t.Logf("DiceRoll(-1, 8) success; expected %v, got %v", "error", err_inv_num)
	}

	// test for an invaild die
	result_inv_die, err_inv_die := DiceRoll(2, 7)

	if result_inv_die != 0 || err_inv_die == nil {
		t.Errorf("DiceRoll(2, 7) failed; expected %v, got %v", "error", result_inv_die)
	} else {
		t.Logf("DiceRoll(2, 7) success; expected %v, got %v", "error", err_inv_die)
	}

	// test for valid arguments
	result_valid, err_valid := DiceRoll(4, 6)

	if err_valid != nil {
		t.Errorf("DiceRoll(4, 6) failed; expected %v, got %v", "4d6", err_valid)
	} else {
		t.Logf("DiceRoll(4, 6) success; expected %v, got %v", "4d6", result_valid)
	}
}

func TestDiceRollVerbose(t *testing.T) {

	// test for invalid number of rolls
	result, err := DiceRollVerbose(0, 20)

	if len(result) != 0 || err == nil {
		t.Errorf("DiceRollVerbose(0, 20) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("DiceRollVerbose(0, 20) success; expected %v, got %v", "error", err)
	}

	// test for an invaild die
	result, err = DiceRollVerbose(1, 3)

	if len(result) != 0 || err == nil {
		t.Errorf("DiceRollVerbose(1, 3) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("DiceRollVerbose(1, 3) success; expected %v, got %v", "error", err)
	}

	// test for valid arguments
	result, err = DiceRollVerbose(4, 6)

	if err != nil {
		t.Errorf("DiceRollVerbose(4, 6) failed; expected %v, got %v", "4d6", err)
	} else {
		t.Logf("DiceRollVerbose(4, 6) success; expected %v, got %v", "4d6", result)
	}
}
