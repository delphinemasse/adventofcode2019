package main

import "testing"

func TestFuelCalc(t *testing.T) {

	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		got := fuelCalc(c.in)
		if got != c.want {
			t.Errorf("fuelCalc(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}