package main

import "testing"

func TestSum(t *testing.T) {
	input1 := 10
	input2 := 20
	expect := 30
	got := sum(input1, input2)

	if got != expect {
		t.Fail()
	}
}
