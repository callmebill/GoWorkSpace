package main

import (
	"strconv"
	"testing"
)

func TestCal(t *testing.T) {
	cases := []struct {
		in1, in2, want string
	}{
		{"2+3", "+", "5"},
		{"2*3", "*", "6"},
		{"3-2", "-", "1"},
	}
	for _, c := range cases {
		got := Cal(c.in1, c.in2)
		want, _ := strconv.Atoi(c.want)
		if got != want {
			t.Errorf("Cal(%q,%q) == %q, want %q", c.in1, c.in2, got, want)
		}
	}
}
