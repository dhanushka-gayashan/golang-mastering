package main

import "testing"

type test struct {
	data []int
	answer int
}

func TestMySum(t *testing.T) {
	tests := []test {
		test{[]int{10, 10}, 20},
		test{[]int{20, 30}, 50},
		test{[]int{50, 40}, 90},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func TestMySumFail(t *testing.T) {
	tests := []test {
		test{[]int{10, 10}, 20},
		test{[]int{20, 30}, 50},
		test{[]int{50, 40}, 90},
	}

	for _, v := range tests {
		x := mySumFail(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}


