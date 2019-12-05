package main

import "testing"

func TestRecursiveFuelCalc(t *testing.T) {

	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		got := recursiveFuelCalc(c.in)
		if got != c.want {
			t.Errorf("recursiveFuelCalc(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}