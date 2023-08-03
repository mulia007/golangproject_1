package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}

	if Addition(-1, -2) != -3 {
		t.Error("Expected (-1) + (-2) to equal -3")
	}
}

func TestMultiple(t *testing.T) {
	if Multiple(2, 2) != 4 {
		t.Error("Expected 2 x 2 to equal 4")
	}

	if Multiple(-1, -2) != 2 {
		t.Error("Expected (-1) * (-2) to equal 2")
	}
}
