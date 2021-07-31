package dice

import "testing"

func TestDice(t *testing.T) {

	// test for invalid number of rolls
	result_inv_num, err_inv_num := dice(-1, 8)

	if result_inv_num != 0 || err_inv_num == nil {
		t.Errorf("dice(-1, 8) failed; expected %v, got %v", "error", result_inv_num)
	} else {
		t.Logf("dice(-1, 8) success; expected %v, got %v", "error", err_inv_num)
	}

	// test for an invaild die
	result_inv_die, err_inv_die := dice(2, 7)

	if result_inv_die != 0 || err_inv_die == nil {
		t.Errorf("dice(2, 7) failed; expected %v, got %v", "error", result_inv_die)
	} else {
		t.Logf("dice(2, 7) success; expected %v, got %v", "error", err_inv_die)
	}

	// test for valid arguments
	result_valid, err_valid := dice(4, 6)

	if err_valid != nil {
		t.Errorf("dice(4, 6) failed; expected %v, got %v", "4d6", err_valid)
	} else {
		t.Logf("dice(4, 6) success; expected %v, got %v", "4d6", result_valid)
	}
}

func TestDiceVerbose(t *testing.T) {

	// test for invalid number of rolls
	result, err := diceVerbose(0, 20)

	if len(result) != 0 || err == nil {
		t.Errorf("diceVerbose(0, 20) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("diceVerbose(0, 20) success; expected %v, got %v", "error", err)
	}

	// test for an invaild die
	result, err = diceVerbose(1, 3)

	if len(result) != 0 || err == nil {
		t.Errorf("diceVerbose(1, 3) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("diceVerbose(1, 3) success; expected %v, got %v", "error", err)
	}

	// test for valid arguments
	result, err = diceVerbose(4, 6)

	if err != nil {
		t.Errorf("diceVerbose(4, 6) failed; expected %v, got %v", "4d6", err)
	} else {
		t.Logf("diceVerbose(4, 6) success; expected %v, got %v", "4d6", result)
	}
}
