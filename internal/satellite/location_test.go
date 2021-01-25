package satellite

import (
	"testing"
)

// deleteElementFromArray
func TestToFixed(t *testing.T) {
	var number = 1.1234567
	var expect1decimalplaces = 1.1
	var expect3decimalplaces = 1.123
	var expect5decimalplaces = 1.12346

	newnumber := toFixed(number, 1)
	if newnumber != expect1decimalplaces {
		t.Errorf("Failure to round number to 1 decimal places: expected %f but got %f",
			expect1decimalplaces, newnumber)
	}

	newnumber = toFixed(number, 3)
	if newnumber != expect3decimalplaces {
		t.Errorf("Failure to round number to 3 decimal places: expected %f but got %f",
			expect3decimalplaces, newnumber)
	}

	newnumber = toFixed(number, 5)
	if newnumber != expect5decimalplaces {
		t.Errorf("Failure to round number to 5 decimal places: expected %f but got %f",
			expect5decimalplaces, newnumber)
	}
}

func TestFindIntersectionBetweenThreeCircles_lessThanThreeCircles(t *testing.T) {
	// the method contract doesn't allow less than 3 cirlces as input
}

func TestFindIntersectionBetweenThreeCircles(t *testing.T) {
	c := [3]circle{
		{
			point: point{
				X: -40.0,
				Y: 0.0,
			},
			R: 40.0,
		},
		{
			point: point{
				X: 0.0,
				Y: 40.0,
			},
			R: 40.0,
		},
		{
			point: point{
				X: -40.0,
				Y: 80.0,
			},
			R: 40.0,
		},
	}

	p, err := findIntersectionBetweenThreeCircles(c)

	if err != nil {
		t.Errorf("Failure to calculate intersection between 3 circles - err")
	}

	expect := point{-40, 40}
	if p != expect {
		t.Errorf("Failure to calculate intersection between 3 circles - point")
	}
}
