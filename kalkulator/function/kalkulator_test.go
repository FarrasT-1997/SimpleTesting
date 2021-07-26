package function

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(2, 4) != 6 {
		t.Error("2 + 4 seharusnya 6")
	}
	if Addition(-2, 4) != 2 {
		t.Error("2 + 4 seharusnya 2")
	}
	if Addition(-2, -4) != -6 {
		t.Error("2 + 4 seharusnya -6")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(2, 4) != -2 {
		t.Error("2 + 4 seharusnya -2")
	}
	if Substraction(-2, 4) != -6 {
		t.Error("2 + 4 seharusnya -6")
	}
	if Substraction(-2, -4) != 2 {
		t.Error("2 + 4 seharusnya 2")
	}
}
func TestDivision(t *testing.T) {
	if Division(2, 4) != 0.5 {
		t.Error("2 + 4 seharusnya 0.5")
	}
	if Division(-2, 4) != -0.5 {
		t.Error("2 + 4 seharusnya -0.5")
	}
	if Division(-2, -4) != 0.5 {
		t.Error("2 + 4 seharusnya 0.5")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 4) != 8 {
		t.Error("2 + 4 seharusnya 8")
	}
	if Multiplication(-2, 4) != -8 {
		t.Error("2 + 4 seharusnya -8")
	}
	if Multiplication(-2, -4) != 8 {
		t.Error("2 + 4 seharusnya 8")
	}
}
