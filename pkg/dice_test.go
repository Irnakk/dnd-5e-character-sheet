package dice

import "testing"

func TestDice(t *testing.T) {

	// test for invalid number of rolls
	result, err := Dice(-1, 8)

	if result != 0 || err == nil {
		t.Errorf("Dice(-1, 8) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("Dice(-1, 8) success; expected %v, got %v", "error", err)
	}

	// test for an invaild die
	result, err = Dice(2, 7)

	if result != 0 || err == nil {
		t.Errorf("Dice(2, 7) failed; expected %v, got %v", "error", result)
	} else {
		t.Logf("Dice(2, 7) success; expected %v, got %v", "error", err)
	}

	// test for valid arguments
	result, err = Dice(4, 6)

	if err != nil {
		t.Errorf("Dice(4, 6) failed; expected %v, got %v", "4d6", err)
	} else {
		t.Logf("Dice(4, 6) success; expected %v, got %v", "4d6", result)
	}
}
