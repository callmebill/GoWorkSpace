package main

import "testing"

func TestCal(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"2+3", "+"},
		{"2*3", "x"},
		{"3-2", "-"},
	}
	for _, c := range cases {
		got := Cal(c.in)
		if got != c.want {
			t.Errorf("Cal(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
